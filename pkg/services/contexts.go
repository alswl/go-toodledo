package services

import (
	"errors"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/context"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/thoas/go-funk"
)

func FindContextByName(auth runtime.ClientAuthInfoWriter, name string) (*models.Context, error) {
	cli := client.NewHTTPClient(strfmt.NewFormats())
	ts, err := cli.Context.GetContextsGetPhp(context.NewGetContextsGetPhpParams(), auth)
	if err != nil {
		return nil, err
	}
	filtered := funk.Filter(ts.Payload, func(x *models.Context) bool {
		return x.Name == name
	}).([]*models.Context)
	if len(filtered) == 0 {
		return nil, errors.New("not found")
	}
	f := filtered[0]
	return f, nil
}
