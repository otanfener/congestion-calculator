package responses

import (
	"github.com/otanfener/congestion-controller/pkg/models"
	"net/http"
)

type CongestionResponse struct {
	Tax models.Tax `json:"tax"`
}

func (CongestionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
