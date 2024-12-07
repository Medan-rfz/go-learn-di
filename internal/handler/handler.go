package handler

import (
	"context"

	"go.uber.org/dig"
)

// IRepo dependency of the repository
type IRepo interface {
	Func(ctx context.Context) error
}

// Handler handler
type Handler struct {
	repo IRepo
}

// Params initialization params
type Params struct {
	dig.In

	Repo IRepo
}

// NewHandler constructor
func NewHandler(p Params) *Handler {
	return &Handler{
		repo: p.Repo,
	}
}
