package shared

import (
	"api/internal/core"
	"context"
	"encoding/json"
	"net/http"
)

type ReactionHandler struct {
	Getter func(ctx context.Context, id string) (*core.Reaction, error)
	Setter func(ctx context.Context, id string, r core.Reaction) error
	Delete func(ctx context.Context, id string) error
}

func (r ReactionHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Path
	switch req.Method {
	case http.MethodGet:
		if r, err := r.Getter(req.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(&r)
		}

	case http.MethodPost:
		d := core.Reaction{}
		if err := json.NewDecoder(req.Body).Decode(&d); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := r.Setter(req.Context(), id, d); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		if err := r.Delete(req.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}
