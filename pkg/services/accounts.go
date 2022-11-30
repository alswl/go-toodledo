package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/account"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	"github.com/pkg/errors"
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
	if err == dal.ErrObjectNotFound {
		me, err := s.Me()
		if err != nil {
			return nil, false, errors.Wrapf(err, "get account s.Me() failed")
		}
		bytes, err = json.Marshal(me)
		if err != nil {
			return nil, false, errors.Wrapf(err, "marshal account failed")
		}
		_ = s.db.Put(BucketAccount, meKey, bytes)
		return me, false, nil
	} else if err != nil {
		return nil, false, errors.Wrapf(err, "get account in db failed")
	} else {
		err = json.Unmarshal(bytes, &u)
		if err != nil {
			return nil, false, errors.Wrapf(err, "unmarshal account failed")
		}
		// TODO check user id for local config
		userID := viper.GetString(models.AuthUserId)
		if u.Userid != userID {
			return nil, false, errors.Errorf("user id not match, local %s, remote %s, please auth logout", userID, u.Userid)
		}

		return &u, true, nil
	}
}

func (s *accountService) GetLastFetchInfo() (*models.Account, error) {
	bytes, err := s.db.Get(BucketAccount, lastSyncInfoKey)
	if err == dal.ErrObjectNotFound {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "get last sync info")
	}
	u := models.Account{}
	err = json.Unmarshal(bytes, &u)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal last sync info")
	}
	return &u, nil
}

func (s *accountService) SetLastFetchInfo(account *models.Account) error {
	bytes, err := json.Marshal(account)
	if err != nil {
		return errors.Wrap(err, "marshal last sync info")
	}
	err = s.db.Put(BucketAccount, lastSyncInfoKey, bytes)
	if err != nil {
		return errors.Wrap(err, "set last sync info")
	}
	return nil
}
