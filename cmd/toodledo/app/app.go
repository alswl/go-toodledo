package app

import (
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/go-openapi/runtime"
)

type ToodledoCliApp struct {
	// TODO service here
	auth    runtime.ClientAuthInfoWriter
	taskSvc services.TaskService
}

func NewToodledoCliApp(auth runtime.ClientAuthInfoWriter, taskSvc services.TaskService) *ToodledoCliApp {
	return &ToodledoCliApp{auth: auth, taskSvc: taskSvc}
}
