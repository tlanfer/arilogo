package twitchhandler

import (
	twitchchat "api/internal/adapter/inbound/twitch"
	"api/internal/core"
	"encoding/json"
	"net/http"
)

func NewHandler(repo core.GlobalConfigRepo, tc twitchchat.Twitch) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			channel, err := repo.GetChannel(req.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			dto := Dto{
				Channel: channel,
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(&dto)

		case http.MethodPost:
			dto := Dto{}
			if err := json.NewDecoder(req.Body).Decode(&dto); err != nil {
				http.Error(w, err.Error(), http.StatusBadGateway)
				return
			}

			if err := repo.SetChannel(req.Context(), dto.Channel); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := tc.Connect(dto.Channel); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})
}

type Dto struct {
	Channel string `json:"channel"`
}
