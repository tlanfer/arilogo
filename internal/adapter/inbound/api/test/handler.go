package test

import (
	"api/internal/core"
	"encoding/json"
	"log"
	"net/http"
)

func NewHandler(wl *core.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "nope", http.StatusMethodNotAllowed)
			return
		}

		dto := Dto{}
		if err := json.NewDecoder(req.Body).Decode(&dto); err != nil {
			log.Printf("failed to parse: %v", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		wl.AddToQueue(core.Reaction{
			PresetId: dto.PresetId,
			Duration: dto.Duration,
		})
	})
}

type Dto struct {
	Duration int           `json:"duration"`
	PresetId core.PresetId `json:"presetId"`
}
