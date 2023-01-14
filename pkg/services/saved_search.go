package services

import (
	"github.com/alswl/go-toodledo/pkg/client0"
	"github.com/alswl/go-toodledo/pkg/client0/saved_search"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/go-openapi/runtime"
	"github.com/sirupsen/logrus"
)

type SavedSearchService interface {
	ListAll() ([]*models.SavedSearch, error)
}

type savedSearchService struct {
	cli  *client0.Toodledo
	auth runtime.ClientAuthInfoWriter
}

func NewSavedSearchService(cli *client0.Toodledo, auth runtime.ClientAuthInfoWriter) SavedSearchService {
	return &savedSearchService{cli: cli, auth: auth}
}

func (s *savedSearchService) ListAll() ([]*models.SavedSearch, error) {
	logrus.Debug("Listing all saved searches")
	p := saved_search.NewGetTasksSearchPhpParams()
	resp, err := s.cli.SavedSearch.GetTasksSearchPhp(p, s.auth)
	if err != nil {
		logrus.WithField("resp", resp).WithError(err).Error("list all saved searches")
		return nil, err
	}
	var list []*models.SavedSearch
	for _, v := range resp.Payload {
		ptr := v
		list = append(list, &ptr)
	}
	return list, nil
}
