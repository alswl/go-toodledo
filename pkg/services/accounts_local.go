package services

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/alswl/go-toodledo/pkg/common"

	"github.com/alswl/go-toodledo/pkg/client0"
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
func CurrentUser(svc AccountExtService) (*models.Account, error) {
	me, _, err := svc.CachedMe()
	return me, err
}

type AccountExtService interface {
	Me() (*models.Account, error)
	CachedMe() (*models.Account, bool, error)
	// GetLastFetchInfo return last fetch info in db
	// lastFetchInfo maybe nil if never sync
	FindLastFetchInfo() (*models.Account, error)
	ModifyLastFetchInfo(account *models.Account) error
}

type accountLocalService struct {
	svc AccountService
	cli *client0.Toodledo
	// TODO access service should now own auth, auth is owned by single user
	auth runtime.ClientAuthInfoWriter
	// TODO no db access
	db dal.Backend
}

func NewAccountExtService(
	cli *client0.Toodledo,
	auth runtime.ClientAuthInfoWriter,
	db dal.Backend,
	accountSvc AccountService,
) AccountExtService {
	return &accountLocalService{cli: cli, auth: auth, db: db, svc: accountSvc}
}

// Me return current user by cli authentication info.
func (s *accountLocalService) Me() (*models.Account, error) {
	return s.svc.Me()
}

func (s *accountLocalService) CachedMe() (*models.Account, bool, error) {
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

func (s *accountLocalService) FindLastFetchInfo() (*models.Account, error) {
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

func (s *accountLocalService) ModifyLastFetchInfo(account *models.Account) error {
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
