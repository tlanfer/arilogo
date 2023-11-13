package byId

import (
	"api/internal/core"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func NewHandler(repo core.ReSubRepo) http.Handler {
	h := handler{
		repo: repo,
	}
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			h.get(w, req)
		case http.MethodPost:
			h.post(w, req)
		case http.MethodDelete:
			h.delete(w, req)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
}

type handler struct {
	repo core.ReSubRepo
}

func (h handler) get(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Path
	all, _ := h.repo.GetAllReSubs(req.Context())
	for _, alert := range all {
		if alert.Id == id {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(Sub{
				Id:       alert.Id,
				Months:   alert.Months,
				State:    alert.State,
				Duration: int64(alert.Duration.Truncate(time.Second).Seconds()),
			})
			return
		}
	}
	http.NotFound(w, req)
}

func (h handler) post(w http.ResponseWriter, req *http.Request) {
	dto := &Sub{}
	if err := json.NewDecoder(req.Body).Decode(dto); err != nil {
		log.Printf("Failed to parse: %v", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.repo.SaveSubAlert(req.Context(), core.ReSub{
		Id:       dto.Id,
		Months:   dto.Months,
		State:    dto.State,
		Duration: time.Duration(dto.Duration) * time.Second,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h handler) delete(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Path

	err := h.repo.DeleteSubAlert(req.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Sub struct {
	Id       string     `json:"id"`
	Months   int        `json:"months"`
	State    core.State `json:"state"`
	Duration int64      `json:"duration"`
}
