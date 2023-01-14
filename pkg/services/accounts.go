package services

import (
	"github.com/alswl/go-toodledo/pkg/client0"
	"github.com/alswl/go-toodledo/pkg/client0/account"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
)

type AccountService interface {
	Me() (*models.Account, error)
}

type accountService struct {
	cli  *client0.Toodledo
	auth runtime.ClientAuthInfoWriter
}

func NewAccountService(cli *client0.Toodledo, auth runtime.ClientAuthInfoWriter) AccountService {
	return &accountService{cli: cli, auth: auth}
}

func (s *accountService) Me() (*models.Account, error) {
	p := account.NewGetAccountGetPhpParams()
	resp, err := s.cli.Account.GetAccountGetPhp(p, s.auth)
	if err != nil {
		return nil, err
	}
	return resp.Payload, nil
}
