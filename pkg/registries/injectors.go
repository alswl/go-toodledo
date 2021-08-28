//go:build wireinject
// +build wireinject

package registries

import (
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/go-openapi/runtime"
	"github.com/google/wire"
)

func InitAuth() (runtime.ClientAuthInfoWriter, error) {
	wire.Build(SuperSet)
	return &auth.SimpleAuth{}, nil
}
