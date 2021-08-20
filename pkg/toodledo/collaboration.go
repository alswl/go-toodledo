package toodledo

import (
	"context"
	"github.com/alswl/go-toodledo/pkg/toodledo/models"
)

// TODO
type CollaboratorService interface {
	Get(ctx context.Context) (models.Collaborators, Response, error)
	ReassignTask(ctx context.Context, id int) (Response, error)
	ShareTask(ctx context.Context, id int) (Response, error)
}

type collaboratorService Service
