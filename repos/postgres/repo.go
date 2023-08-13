package postgres

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/otanfener/congestion-controller/pkg/db/postgres"
	"github.com/otanfener/congestion-controller/pkg/models"
	"github.com/otanfener/congestion-controller/repos"
)

var _ repos.Repo = &Repo{}

type Repo struct {
	db      *postgres.Service
	builder squirrel.StatementBuilderType
}

func (r *Repo) GetCity(ctx context.Context, city string) (models.City, error) {
	return models.City{}, nil
}
