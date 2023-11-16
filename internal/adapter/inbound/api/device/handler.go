package device

import (
	"api/internal/core"
	"encoding/json"
	"net/http"
	"strings"
)

func NewDeviceHandler(config core.GlobalConfigRepo, light core.Light) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			addr, err := config.GetDeviceAddr(req.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			dto := Dto{Address: strings.TrimSpace(addr)}
			_ = json.NewEncoder(w).Encode(&dto)

		case http.MethodPost:
			dto := Dto{}
			if err := json.NewDecoder(req.Body).Decode(&dto); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			addr := strings.TrimSpace(dto.Address)

			if err := config.SetDeviceAddr(req.Context(), addr); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			light.SetAddr(addr)
		}
	})
}

type Dto struct {
	Address string `json:"address"`
}
