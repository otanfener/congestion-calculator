package responses

import (
	"github.com/go-chi/render"
	"github.com/otanfener/congestion-controller/pkg/domain"
	"net/http"
)

type ErrorResponse struct {
	HTTPCode int   `json:"-"`
	Error    error `json:"error"`
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPCode)
	return nil
}

func ErrInternal() render.Renderer {
	return ErrorResponse{
		HTTPCode: http.StatusInternalServerError,
		Error:    domain.ErrInternal,
	}
}

func ErrBadRequest() render.Renderer {
	return ErrorResponse{
		HTTPCode: http.StatusBadRequest,
		Error:    domain.ErrBadRequest,
	}
}

func ErrNotFound() render.Renderer {
	return ErrorResponse{
		HTTPCode: http.StatusNotFound,
		Error:    domain.ErrNotFound,
	}
}
