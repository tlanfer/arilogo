package patternById

import (
	"api/internal/core"
	"encoding/json"
	"log"
	"net/http"
)

func NewHandler(repo core.PatternRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			id := req.URL.Path
			all, _ := repo.GetAllPattern(req.Context())
			for _, pattern := range all {
				if pattern.Id == id {
					w.Header().Set("Content-Type", "application/json")
					_ = json.NewEncoder(w).Encode(pattern)
					return
				}
			}
			http.NotFound(w, req)
		case http.MethodPost:
			dto := core.Pattern{}
			if err := json.NewDecoder(req.Body).Decode(&dto); err != nil {
				log.Printf("Failed to parse: %v", err.Error())
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := repo.SavePattern(req.Context(), dto); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})

}
