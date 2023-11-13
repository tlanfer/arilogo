package test

import (
	"api/internal/core"
	"api/internal/core/controller"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func NewHandler(wl *controller.Controller) http.Handler {
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

		wl.AddToQueue(core.Action{
			State:    dto.State,
			Duration: time.Duration(dto.Duration) * time.Second,
		})
	})
}

type Dto struct {
	Duration int        `json:"duration"`
	State    core.State `json:"state"`
}
