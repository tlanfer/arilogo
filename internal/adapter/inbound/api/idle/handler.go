package idle

import (
	"api/internal/core"
	"encoding/json"
	"log"
	"net/http"
)

func NewHandler(repo core.GlobalConfigRepo, c *core.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			presetId, err := repo.GetIdleState(req.Context())
			if err != nil {
				log.Printf("failed to load: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			dto := Dto{
				IdlePresetId: presetId,
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(&dto); err != nil {
				log.Printf("failed to serialize response: %v", err)
			}
		case http.MethodPost:
			dto := Dto{}
			if err := json.NewDecoder(req.Body).Decode(&dto); err != nil {
				log.Printf("failed to parse request: %v", err)
			}

			if err := repo.SetIdleState(req.Context(), dto.IdlePresetId); err != nil {
				log.Printf("failed to save: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			c.ConfigChanged()
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
}

type Dto struct {
	IdlePresetId core.PresetId `json:"idlePresetId"`
}
