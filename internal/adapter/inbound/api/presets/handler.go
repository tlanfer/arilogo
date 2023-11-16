package presets

import (
	"api/internal/core"
	"encoding/json"
	"log"
	"net/http"
)

func NewHandler(c core.Light) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			presets := c.Presets()
			dto := Dto{
				Presets: make([]PresetDto, 0),
			}
			for id, name := range presets {
				if name == "" {
					continue
				}
				dto.Presets = append(dto.Presets, PresetDto{
					Id:   int(id),
					Name: name,
				})
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(dto); err != nil {
				log.Printf("failed to serialize presets: %v", err)
			}
		}
	})
}

type Dto struct {
	Presets []PresetDto `json:"presets"`
}

type PresetDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
