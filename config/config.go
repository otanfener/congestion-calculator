package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/otanfener/congestion-controller/pkg/db/mongo"
)

type Config struct {
	Addr string       `required:"true"`
	DB   mongo.Config `envconfig:"db"`
}

func New() (Config, error) {
	var cfg Config
	err := envconfig.Process("volvo", &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, err
}
