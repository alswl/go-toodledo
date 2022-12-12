package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/alswl/go-toodledo/pkg/common"

	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/account"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	"github.com/spf13/viper"
)

const BucketAccount = "account"
const lastSyncInfoKey = "lastSyncInfo"

// meKey is the key of account in db, TODO multi user support
const meKey = "me"

// CurrentUser return live user
// TODO move to cli
func CurrentUser(svc AccountService) (*models.Account, error) {
	me, _, err := svc.CachedMe()
	return me, err
}

// AccountService ...
// TODO split to AccountService and SyncService
type AccountService interface {
	Me() (*models.Account, error)
	CachedMe() (*models.Account, bool, error)
	// GetLastFetchInfo return last fetch info in db
	// lastFetchInfo maybe nil if never sync
	GetLastFetchInfo() (*models.Account, error)
	SetLastFetchInfo(account *models.Account) error
}

type accountService struct {
	cli  *client.Toodledo
	auth runtime.ClientAuthInfoWriter
	db   dal.Backend
}

// NewAccountService ...
func NewAccountService(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter, db dal.Backend) AccountService {
	return &accountService{cli: cli, auth: auth, db: db}
}

func (s *accountService) Me() (*models.Account, error) {
	p := account.NewGetAccountGetPhpParams()
	resp, err := s.cli.Account.GetAccountGetPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

func (s *accountService) CachedMe() (*models.Account, bool, error) {
	bytes, err := s.db.Get(BucketAccount, meKey)
	u := models.Account{}
	// TODO refactor, cacheOrSet(key, setFn): val

	if err != nil {
		if errors.Is(err, dal.ErrObjectNotFound) {
			me, ierr := s.Me()
			if ierr != nil {
				return nil, false, fmt.Errorf("get account s.Me() failed: %w", ierr)
			}
			bytes, ierr = json.Marshal(me)
			if ierr != nil {
				return nil, false, fmt.Errorf("marshal account failed: %w", ierr)
			}
			_ = s.db.Put(BucketAccount, meKey, bytes)
			return me, false, nil
		}
		return nil, false, fmt.Errorf("get account in db failed: %w", err)
	}
	err = json.Unmarshal(bytes, &u)
	if err != nil {
		return nil, false, fmt.Errorf("unmarshal account failed: %w", err)
	}
	// TODO check user id for local config
	userID := viper.GetString(common.AuthUserID)
	if u.Userid != userID {
		return nil, false, fmt.Errorf("user id not match, local %s, remote %s, please auth logout", userID, u.Userid)
	}

	return &u, true, nil
}

func (s *accountService) GetLastFetchInfo() (*models.Account, error) {
	bytes, err := s.db.Get(BucketAccount, lastSyncInfoKey)
	if errors.Is(err, dal.ErrObjectNotFound) {
		return nil, common.ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("get last sync info failed: %w", err)
	}
	u := models.Account{}
	err = json.Unmarshal(bytes, &u)
	if err != nil {
		return nil, fmt.Errorf("unmarshal last sync info failed: %w", err)
	}
	return &u, nil
}

func (s *accountService) SetLastFetchInfo(account *models.Account) error {
	bytes, err := json.Marshal(account)
	if err != nil {
		return fmt.Errorf("marshal last sync info failed: %w", err)
	}
	err = s.db.Put(BucketAccount, lastSyncInfoKey, bytes)
	if err != nil {
		return fmt.Errorf("set last sync info failed: %w", err)
	}
	return nil
}
