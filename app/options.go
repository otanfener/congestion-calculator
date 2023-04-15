package app

type Option func(*API)

func WithCongestionSrv(srv CongestionService) Option {
	return func(api *API) {
		api.congestionService = srv
	}
}
