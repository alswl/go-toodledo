package app

import (
	"github.com/alswl/go-toodledo/pkg/models"
)

type ToodledoCliApp struct {
	CurrentUser *models.Account
}

func NewToodledoCliApp(currentUser *models.Account) *ToodledoCliApp {
	return &ToodledoCliApp{CurrentUser: currentUser}
}
