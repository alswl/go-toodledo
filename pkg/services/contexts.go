package services

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/context"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// ContextService ...
type ContextService interface {
	Find(name string) (*models.Context, error)
	FindByID(id int64) (*models.Context, error)
	ListAll() ([]*models.Context, error)
	Rename(name string, newName string) (*models.Context, error)
	Delete(name string) error
	Create(name string) (*models.Context, error)
}

// ContextPersistenceService is a cached service
// it synced interval by fetcher.
type ContextPersistenceService interface {
	Synchronizable
	ContextService
}

type contextService struct {
	cli  *client.Toodledo
	auth runtime.ClientAuthInfoWriter
}

// NewContextService ...
func NewContextService(cli *client.Toodledo, auth runtime.ClientAuthInfoWriter) ContextService {
	return &contextService{cli: cli, auth: auth}
}

// Create ...
func (s *contextService) Create(name string) (*models.Context, error) {
	params := context.NewPostContextsAddPhpParams()
	params.SetName(name)
	resp, err := s.cli.Context.PostContextsAddPhp(params, s.auth)
	if err != nil {
		logrus.WithField("resp", resp).WithError(err).Error("create")
		return nil, err
	}
	return resp.Payload[0], nil
}

// Delete ...
func (s *contextService) Delete(name string) error {
	f, err := s.Find(name)
	if err != nil {
		return err
	}

	params := context.NewPostContextsDeletePhpParams()
	params.SetID(f.ID)
	resp, err := s.cli.Context.PostContextsDeletePhp(params, s.auth)
	if err != nil {
		logrus.WithField("resp", resp).WithError(err).Error("delete context")
		return err
	}
	return nil
}

// Rename ...
func (s *contextService) Rename(name string, newName string) (*models.Context, error) {
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

// Find ...
func (s *contextService) Find(name string) (*models.Context, error) {
	logrus.Warn("FindByID is implemented with ListALl(), it's deprecated, please using cache")
	fs, err := s.ListAll()
	if err != nil {
		return nil, err
	}

	filtered, _ := funk.Filter(fs, func(x *models.Context) bool {
		return x.Name == name
	}).([]*models.Context)
	if len(filtered) == 0 {
		return nil, common.ErrNotFound
	}
	f := filtered[0]
	return f, nil
}

func (s *contextService) FindByID(id int64) (*models.Context, error) {
	logrus.Warn("FindByID is implemented with ListALl(), it's deprecated, please using cache")
	fs, err := s.ListAll()
	if err != nil {
		return nil, err
	}

	filtered, _ := funk.Filter(fs, func(x *models.Context) bool {
		return x.ID == id
	}).([]*models.Context)
	if len(filtered) == 0 {
		return nil, common.ErrNotFound
	}
	f := filtered[0]
	return f, nil
}

// ListAll ...
func (s *contextService) ListAll() ([]*models.Context, error) {
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
	filtered, _ := funk.Filter(ts.Payload, func(x *models.Context) bool {
		return x.Name == name
	}).([]*models.Context)
	if len(filtered) == 0 {
		return nil, errors.New("not found")
	}
	f := filtered[0]
	return f, nil
}
