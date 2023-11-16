package streamlabshandler

import (
	"api/internal/core"
	"encoding/json"
	"github.com/tlanfer/go-streamlabs"
	"net/http"
)

func NewHandler(config core.GlobalConfigRepo, sl streamlabs.Streamlabs) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			token, err := config.GetStreamlabsToken(req.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if token != "" {
				token = "redactedredactedredactedredactedredacted"
			}
			dto := Dto{Token: token}
			_ = json.NewEncoder(w).Encode(&dto)

		case http.MethodPost:
			dto := Dto{}
			if err := json.NewDecoder(req.Body).Decode(&dto); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := config.SetStreamlabsToken(req.Context(), dto.Token); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			sl.Connect(dto.Token)
		}
	})
}

type Dto struct {
	Token string `json:"token"`
}
