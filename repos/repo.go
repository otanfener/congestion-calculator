package repos

import (
	"context"
	"fmt"
	"github.com/otanfener/congestion-controller/pkg/db"
	"github.com/otanfener/congestion-controller/pkg/domain"
	"github.com/otanfener/congestion-controller/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	db         *db.DB
	collection string
}

func New(database *db.DB, collection string) *Repo {
	repo := &Repo{db: database, collection: collection}
	return repo
}

func (r *Repo) GetCity(ctx context.Context, city string) (models.City, error) {
	collection := r.db.Collection(r.collection)
	fmt.Printf("received city:%s\n", city)
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
