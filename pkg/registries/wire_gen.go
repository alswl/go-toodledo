// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package registries

import (
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/go-openapi/runtime"
)

// Injectors from injectors.go:

func InitAuth() (runtime.ClientAuthInfoWriter, error) {
	string2, err := auth.ProvideAccessToken()
	if err != nil {
		return nil, err
	}
	clientAuthInfoWriter, err := auth.ProvideSimpleAuth(string2)
	if err != nil {
		return nil, err
	}
	return clientAuthInfoWriter, nil
}
