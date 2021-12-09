package toodledo

import (
	"context"
	"github.com/alswl/go-toodledo/pkg/toodledo/models"
)

// AccountService ...
type AccountService Service

// Get ...
func (s *AccountService) Get(ctx context.Context) (*models.Account, *Response, error) {
	path := "/3/account/get.php"

	req, err := s.client.NewRequest("GET", path)
	if err != nil {
		return nil, nil, err
	}

	var account *models.Account
	resp, err := s.client.Do(ctx, req, &account)
	if err != nil {
		return nil, resp, err
	}

	return account, resp, nil
}
