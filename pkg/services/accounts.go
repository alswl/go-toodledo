package services

import (
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/account"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
)

// CurrentUser ...
func CurrentUser(svc AccountService) (*models.Account, error) {
	return svc.Me()
}

// AccountService ...
type AccountService interface {
	Me() (*models.Account, error)
}

type accountService struct {
	cli  *client.Toodledo
	auth runtime.ClientAuthInfoWriter
}

// NewAccountService ...
func NewAccountService(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter) AccountService {
	return &accountService{cli: cli, auth: auth}
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
