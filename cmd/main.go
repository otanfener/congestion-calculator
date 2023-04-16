package main

import (
	"fmt"
	"github.com/otanfener/congestion-controller/app"
	"github.com/otanfener/congestion-controller/config"
	"github.com/otanfener/congestion-controller/pkg/db"
	"github.com/otanfener/congestion-controller/repos"
	"github.com/otanfener/congestion-controller/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	errorChannel := make(chan error)
	go func() {
		zlog.Info().Msgf("starting server on %s", cfg.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zlog.Error().Msg("failed to create http server")
			errorChannel <- fmt.Errorf("failed to create http server: %s", err)
		}
	}()
	// Capture interrupts.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errorChannel <- fmt.Errorf("got signal: %s", <-c)
	}()

	// Wait for any error.
	if err := <-errorChannel; err != nil {
		zlog.Error().Msgf("received error: %s", err)
		os.Exit(1)
	}
}
