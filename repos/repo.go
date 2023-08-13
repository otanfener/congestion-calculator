package repos

import (
	"context"
	"github.com/otanfener/congestion-controller/pkg/models"
)

//go:generate moq -out repo_mock.go . Repo
type Repo interface {
	GetCity(ctx context.Context, city string) (models.City, error)
}
