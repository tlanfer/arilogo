package patternList

import (
	"api/internal/core"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func NewHandler(repo core.PatternRepo) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			pattern, err := repo.GetAllPattern(req.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			dto := Dto{}
			for _, p := range pattern {
				dto.Pattern = append(dto.Pattern, p)
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(&dto); err != nil {
				log.Printf("failed to serialize: %v", err.Error())
			}

		case http.MethodPost:
			pattern, _ := repo.GetAllPattern(req.Context())
			dto := core.Pattern{
				Id:   uuid.NewString(),
				Name: fmt.Sprintf("Logo #%v", len(pattern)+1),
				Segments: []core.Segment{
					{Start: 0, End: 70, Color1: core.Color{R: 255, G: 255, B: 255}},
					{Start: 71, End: 150, Color1: core.Color{R: 255, G: 0, B: 0}},
				}}

			if err := repo.SavePattern(req.Context(), dto); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(&dto); err != nil {
				log.Printf("failed to serialize: %v", err.Error())
			}
		}
	})
}

type Dto struct {
	Pattern []core.Pattern `json:"pattern"`
}
