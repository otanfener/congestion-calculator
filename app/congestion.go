package app

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/render"
	"github.com/otanfener/congestion-controller/pkg/domain"
	"github.com/otanfener/congestion-controller/pkg/models"
	"github.com/otanfener/congestion-controller/responses"
	"net/http"
)

type CongestionService interface {
	CalculateTax(ctx context.Context, times []models.CivilTime, city string, car string) (models.Tax, error)
}
type CongestionTaxParams struct {
	Body models.NewCongestionTax
}

func (api *API) CalculateTax() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()

		ctx := request.Context()

		var params CongestionTaxParams

		err := json.NewDecoder(request.Body).Decode(&params.Body)
		if err != nil {
			render.Render(writer, request, responses.ErrBadRequest())
			return
		}
		err = api.validator.Struct(params)
		if err != nil {
			render.Render(writer, request, responses.ErrBadRequest())
			return
		}

		res, err := api.congestionService.CalculateTax(ctx, params.Body.Times, params.Body.City, params.Body.VehicleType)

		if errors.Is(err, domain.ErrNotFound) {
			render.Render(writer, request, responses.ErrNotFound())
			return
		} else if err != nil {
			render.Render(writer, request, responses.ErrInternal())
			return
		}
		resp := responses.CongestionResponse{Tax: res}
		render.Render(writer, request, resp)
		return
	}
}
