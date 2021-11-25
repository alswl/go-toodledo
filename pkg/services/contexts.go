package services

import (
	"errors"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/context"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"strconv"
)

type ContextService interface {
	Find(name string) (*models.Context, error)
	ListAll() ([]*models.Context, error)
	Rename(name string, newName string) (*models.Context, error)
	Delete(name string) error
	Create(name string) (*models.Context, error)
}

type contextservice struct {
	cli  *client.Toodledo
	auth runtime.ClientAuthInfoWriter
}

func NewContextService(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter) ContextService {
	return &contextservice{cli: cli, auth: auth}
}

func (s *contextservice) Create(name string) (*models.Context, error) {
	params := context.NewPostContextsAddPhpParams()
	params.SetName(name)
	resp, err := s.cli.Context.PostContextsAddPhp(params, s.auth)
	if err != nil {
		logrus.WithField("resp", resp).WithError(err).Error("Failed to create")
		return nil, err
	}
	return resp.Payload[0], nil
}

func (s *contextservice) Delete(name string) error {
	f, err := s.Find(name)
	if err != nil {
		return err
	}

	params := context.NewPostContextsDeletePhpParams()
	params.SetID(f.ID)
	resp, err := s.cli.Context.PostContextsDeletePhp(params, s.auth)
	if err != nil {
		logrus.WithField("resp", resp).WithError(err).Error("Failed to delete context")
		return err
	}
	return nil
}

func (s *contextservice) Rename(name string, newName string) (*models.Context, error) {
	if name == newName {
		logrus.Error("not changed")
		return nil, fmt.Errorf("not changed")
	}

	f, err := s.Find(name)
	if err != nil {
		logrus.Error(err)
		return nil, common.ErrNotFound
	}

	p := context.NewPostContextsEditPhpParams()
	p.SetID(strconv.Itoa(int(f.ID)))
	p.SetName(&newName)
	resp, err := s.cli.Context.PostContextsEditPhp(p, s.auth)
	if err != nil {
		logrus.WithField("resp", resp).WithError(err).Error("request failed")
		return nil, err
	}
	return resp.Payload[0], nil
}

func (s *contextservice) Find(name string) (*models.Context, error) {
	fs, err := s.ListAll()
	if err != nil {
		return nil, err
	}

	filtered := funk.Filter(fs, func(x *models.Context) bool {
		return x.Name == name
	}).([]*models.Context)
	if len(filtered) == 0 {
		return nil, common.ErrNotFound
	}
	f := filtered[0]
	return f, nil
}

func (s *contextservice) ListAll() ([]*models.Context, error) {
	cli := client.NewHTTPClient(strfmt.NewFormats())
	ts, err := cli.Context.GetContextsGetPhp(context.NewGetContextsGetPhpParams(), s.auth)
	if err != nil {
		return nil, err
	}
	return ts.Payload, nil
}

// TODO using wire
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
