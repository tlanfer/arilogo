package donations

import (
	"api/internal/adapter/inbound/api/shared"
	"api/internal/core"
	"context"
	"net/http"
)

func ById(repo core.DonationAlertRepo) http.Handler {
	getter := func(ctx context.Context, id string) (*core.Reaction, error) {
		return repo.GetDonationAlertById(ctx, id)
	}
	setter := func(ctx context.Context, id string, r core.Reaction) error {
		return repo.SetDonationAlertById(ctx, id, r)
	}
	deleter := func(ctx context.Context, id string) error { return repo.DeleteDonationAlertById(ctx, id) }

	return shared.ReactionHandler{
		Getter: getter,
		Setter: setter,
		Delete: deleter,
	}
}

func List(repo core.DonationAlertRepo) http.Handler {
	list := func(ctx context.Context) (map[string]core.Reaction, error) {
		return repo.GetAllDonationAlerts(ctx)
	}
	setter := func(ctx context.Context, id string, r core.Reaction) error {
		return repo.SetDonationAlertById(ctx, id, r)
	}

	return shared.ReactionList{
		List:   list,
		Setter: setter,
		Step:   100,
	}
}
