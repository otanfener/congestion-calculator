package mongo

import (
	"context"
	mongodb "github.com/otanfener/congestion-controller/pkg/db/mongo"
	"github.com/otanfener/congestion-controller/pkg/domain"
	"github.com/otanfener/congestion-controller/pkg/models"
	"github.com/otanfener/congestion-controller/repos"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ repos.Repo = &Repo{}

type Repo struct {
	db         *mongodb.Service
	collection string
}

func New(database *mongodb.Service, collection string) *Repo {
	repo := &Repo{db: database, collection: collection}
	return repo
}

func (r *Repo) GetCity(ctx context.Context, city string) (models.City, error) {
	collection := r.db.Collection(r.collection)
	res := collection.FindOne(ctx, bson.M{"name": city})
	switch {
	case res.Err() == mongo.ErrNoDocuments:
		return models.City{}, domain.ErrNotFound
	case res.Err() != nil:
		return models.City{}, res.Err()
	}
	var c models.City
	err := res.Decode(&c)
	if err != nil {
		return models.City{}, domain.ErrInternal
	}
	return c, nil
}
