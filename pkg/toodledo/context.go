package toodledo

import (
	"context"
	"github.com/alswl/go-toodledo/pkg/toodledo/models"
)

// TODO @alswl impl
type ContextService interface {
	Get(ctx context.Context) ([]models.Context, Response, error)
	Add(ctx context.Context, name string) (models.Context, Response, error)
	Edit(ctx context.Context, id int, name string, private bool) (models.Context, Response, error)
	Delete(ctx context.Context, id int) (Response, error)
}

type contextService Service
