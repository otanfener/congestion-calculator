package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	*mongo.Database
}

// mongodb://congestion:congestion@localhost:27017/volvo
func New(cfg Config) (*DB, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	//ping the database
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(cfg.Name)
	return &DB{db}, nil
}
