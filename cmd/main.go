package main

import (
	"github.com/otanfener/congestion-controller/app"
	"github.com/otanfener/congestion-controller/config"
	"github.com/otanfener/congestion-controller/pkg/db"
	"github.com/otanfener/congestion-controller/repos"
	"github.com/otanfener/congestion-controller/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

const (
	defaultReadTimeout  = 5 * time.Second
	defaultWriteTimeout = 5 * time.Second
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	cfg, err := config.New()
	if err != nil {
		log.Error().Msg("failed to parse config file")
		return
	}
	zlog := zerolog.New(os.Stdout)
	database, err := db.New(cfg.DB)
	if err != nil {
		log.Error().Msg("failed to create db connection")
		return
	}

	r := repos.New(database, cfg.DB.Collection)

	var opts []app.Option
	{
		srv := service.New(r)
		opts = append(opts, app.WithCongestionSrv(srv))
	}
	srv := http.Server{
		Addr:         cfg.Addr,
		Handler:      app.New(cfg, zlog, opts...),
		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}
	zlog.Info().Msgf("starting server on %s", cfg.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		zlog.Error().Msg("failed to create http server")
	}
}
