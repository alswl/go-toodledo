package services

import (
	"encoding/json"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/account"
	"github.com/alswl/go-toodledo/pkg/dal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	"github.com/pkg/errors"
)

// CurrentUser ...
func CurrentUser(svc AccountService) (*models.Account, error) {
	return svc.Me()
}

var BucketAccount = "account"

// AccountService ...
type AccountService interface {
	Me() (*models.Account, error)
	CachedMe() (*models.Account, bool, error)
	GetLastSyncInfo() (*models.Account, error)
	SetLastSyncInfo(account *models.Account) error
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

// Me ...
func (s *accountService) Me() (*models.Account, error) {
	p := account.NewGetAccountGetPhpParams()
	resp, err := s.cli.Account.GetAccountGetPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}

var key = "lastSyncInfo"

func (s *accountService) GetLastSyncInfo() (*models.Account, error) {
	bytes, err := s.db.Get(BucketAccount, key)
	if err == dal.ErrObjectNotFound {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "failed to get last sync info")
	}
	u := models.Account{}
	err = json.Unmarshal(bytes, &u)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal last sync info")
	}
	return &u, nil
}

func (s *accountService) SetLastSyncInfo(account *models.Account) error {
	bytes, err := json.Marshal(account)
	if err != nil {
		return errors.Wrap(err, "failed to marshal last sync info")
	}
	err = s.db.Put(BucketAccount, key, bytes)
	if err != nil {
		return errors.Wrap(err, "failed to set last sync info")
	}
	return nil
}

func (s *accountService) cachedMe(key string) (*models.Account, bool, error) {
	bytes, err := s.db.Get(BucketAccount, key)
	u := models.Account{}
	if err == dal.ErrObjectNotFound {
		me, err := s.Me()
		if err != nil {
			return nil, false, errors.Wrapf(err, "get account s.Me() failed")
		}
		bytes, err = json.Marshal(me)
		if err != nil {
			return nil, false, errors.Wrapf(err, "marshal account failed")
		}
		_ = s.db.Put(BucketAccount, key, bytes)
		return me, false, nil
	} else if err != nil {
		return nil, false, errors.Wrapf(err, "get account in db failed")
	} else {
		err = json.Unmarshal(bytes, &u)
		if err != nil {
			return nil, false, errors.Wrapf(err, "unmarshal account failed")
		}
		return &u, true, nil
	}
}

func (s *accountService) CachedMe() (*models.Account, bool, error) {
	return s.cachedMe("me")
}
