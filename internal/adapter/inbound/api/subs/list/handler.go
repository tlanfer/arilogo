package list

import (
	"api/internal/adapter/outbound/wled"
	"api/internal/core"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"sort"
	"time"
)

type repo interface {
	core.ReSubRepo
}

func NewHandler(repo repo, wl *wled.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			allReSubs, err := repo.GetAllReSubs(req.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			dto := Dto{}
			for _, sub := range allReSubs {
				dto.Resubs = append(dto.Resubs, ReSub{
					Id:       sub.Id,
					Months:   sub.Months,
					State:    sub.State,
					Duration: int(sub.Duration.Seconds()),
				})
			}

			if err := json.NewEncoder(w).Encode(&dto); err != nil {
				log.Printf("failed to serialize: %v", err.Error())
			}

		case http.MethodPost:
			allReSubs, err := repo.GetAllReSubs(req.Context())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			n := 0

			if len(allReSubs) > 0 {
				sort.Sort(allReSubs)
				n = allReSubs[len(allReSubs)-1].Months + 1
			}

			dto := ReSub{
				Id:       uuid.NewString(),
				Months:   n,
				Duration: 3,
			}

			err = repo.SaveSubAlert(req.Context(), core.ReSub{
				Id:       dto.Id,
				Months:   dto.Months,
				State:    dto.State,
				Duration: time.Duration(dto.Duration) * time.Second,
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := json.NewEncoder(w).Encode(&dto); err != nil {
				log.Printf("failed to serialize: %v", err.Error())
			}

		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
}

type Dto struct {
	Resubs []ReSub `json:"resubs"`
}

type ReSub struct {
	Id       string     `json:"id"`
	Months   int        `json:"months"`
	Duration int        `json:"duration"`
	State    core.State `json:"state"`
}
