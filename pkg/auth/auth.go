package auth

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

type SimpleAuth struct {
	accessToken string
}

func NewSimpleAuth(accessToken string) *SimpleAuth {
	return &SimpleAuth{accessToken: accessToken}
}

func (a *SimpleAuth) AuthenticateRequest(request runtime.ClientRequest, registry strfmt.Registry) error {
	request.SetQueryParam("access_token", a.accessToken)
	return nil
}
