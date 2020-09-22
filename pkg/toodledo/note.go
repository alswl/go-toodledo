package toodledo

import (
	"context"
	"github.com/alswl/go-toodledo/pkg/toodledo/models"
)

// TODO
type NoteService interface {
	Get(ctx context.Context) ([]models.Note, Response, error)
	Add(ctx context.Context) (models.Note, Response, error)
	Edit(ctx context.Context) (models.Note, Response, error)
	Delete(ctx context.Context) (Response, error)
}

type noteService Service
