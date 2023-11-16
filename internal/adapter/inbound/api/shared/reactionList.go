package shared

import (
	"api/internal/core"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type ReactionList struct {
	List   func(ctx context.Context) (map[string]core.Reaction, error)
	Setter func(ctx context.Context, id string, r core.Reaction) error
	Step   int
}

func (r ReactionList) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		reactions, err := r.List(req.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(Dto{Reactions: reactions})

	case http.MethodPost:

		reactions, err := r.List(req.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		max := 0
		for _, reaction := range reactions {
			if reaction.Amount > max {
				max = reaction.Amount
			}
		}
		max += r.Step
		id := uuid.NewString()
		reaction := core.Reaction{
			Amount:   max,
			PresetId: 1,
			Duration: 2500,
		}

		if err = r.Setter(req.Context(), id, reaction); err != nil {
			log.Printf("failed: %v", err)
		}

		w.Header().Set("Location", id)
		w.WriteHeader(http.StatusCreated)
	}
}

type Dto struct {
	Reactions map[string]core.Reaction `json:"reactions"`
}
