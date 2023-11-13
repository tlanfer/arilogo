package effects

import (
	"api/internal/adapter/outbound/wled"
	"encoding/json"
	"net/http"
)

func NewHandler(wl *wled.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		effects := wl.Effects()
		dto := Dto{Effects: effects}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(&dto)
	})
}

type Dto struct {
	Effects []wled.Effect `json:"effects"`
}
