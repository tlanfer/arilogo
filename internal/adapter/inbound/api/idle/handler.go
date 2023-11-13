package idle

import (
	"api/internal/core"
	"api/internal/core/controller"
	"encoding/json"
	"log"
	"net/http"
)

func NewHandler(repo core.GlobalConfigRepo, c *controller.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			state, err := repo.GetIdleState(req.Context())
			if err != nil {
				log.Printf("failed to load: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			dto := Dto{
				IdleState: *state,
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

			log.Printf("new idle state: %+v", dto.IdleState)
			if err := repo.SetIdleState(req.Context(), dto.IdleState); err != nil {
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
	IdleState core.State `json:"IdleState"`
}
