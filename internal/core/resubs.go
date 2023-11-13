package core

import (
	"context"
	"time"
)

type ReSubRepo interface {
	GetAllReSubs(ctx context.Context) (ReSubs, error)
	SaveSubAlert(ctx context.Context, alert ReSub) error
	DeleteSubAlert(ctx context.Context, id string) error
}

type ReSub struct {
	Id       string
	Months   int
	State    State
	Duration time.Duration
}

type ReSubs []ReSub

func (r ReSubs) Len() int {
	return len(r)
}

func (r ReSubs) Less(i, j int) bool {
	return r[i].Months < r[j].Months
}

func (r ReSubs) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
