package main

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/folder"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type Auth struct {
	accessToken string
}

func NewAuth(accessToken string) *Auth {
	return &Auth{accessToken: accessToken}
}

func (a *Auth) AuthenticateRequest(request runtime.ClientRequest, registry strfmt.Registry) error {
	request.SetQueryParam("access_token", a.accessToken)
	return nil
}

func main() {
	accessToken := os.Getenv("TOODLEDO_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("Unauthorized: No TOODLEDO_ACCESS_TOKEN present")
	}
	auth := NewAuth(accessToken)

	cli := client.NewHTTPClient(strfmt.NewFormats())
	fs, err := cli.Folder.GetFoldersGetPhp(folder.NewGetFoldersGetPhpParams(), auth)
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Printf("%v", fs)
}
