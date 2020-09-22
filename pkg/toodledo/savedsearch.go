package toodledo

import (
	"context"
	"github.com/alswl/go-toodledo/pkg/toodledo/models"
)

// TODO
type SavedSearchService interface {
	Get(ctx context.Context) (models.SavedSearch, Response, error)
}

type savedSearchService Service
