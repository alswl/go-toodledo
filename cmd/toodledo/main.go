package main

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/folder"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
	"os"
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

func main() {
	accessToken := os.Getenv("TOODLEDO_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("Unauthorized: No TOODLEDO_ACCESS_TOKEN present")
	}
	auth := NewSimpleAuth(accessToken)

	cli := client.NewHTTPClient(strfmt.NewFormats())
	res, err := cli.Folder.GetFoldersGetPhp(folder.NewGetFoldersGetPhpParams(), auth)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Print(render.TablesRender(res.Payload))
}
