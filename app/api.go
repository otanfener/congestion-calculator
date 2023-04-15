package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/otanfener/congestion-controller/config"
	"github.com/rs/zerolog"
	"net/http"
)

type API struct {
	cfg               config.Config
	router            chi.Router
	validator         *validator.Validate
	congestionService CongestionService
	logger            zerolog.Logger
}

func New(cfg config.Config, logger zerolog.Logger, opts ...Option) *API {
	app := &API{
		cfg:       cfg,
		logger:    logger,
		validator: validator.New(),
	}

	for _, opt := range opts {
		opt(app)
	}
	app.initRouter()
	return app
}

func (api *API) initRouter() {
	router := chi.NewRouter()

	router.Route("/api", func(r chi.Router) {
		r.Route("/congestion", func(r chi.Router) {
			r.Post("/", api.CalculateTax())
		})
	})

	api.router = router
}

func (api *API) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	api.router.ServeHTTP(writer, request)
}
