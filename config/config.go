package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/otanfener/congestion-controller/pkg/db"
)

type Config struct {
	Addr string    `required:"true"`
	DB   db.Config `envconfig:"db"`
}

func New() (Config, error) {
	var cfg Config
	err := envconfig.Process("volvo", &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, err
}
