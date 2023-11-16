package bitsHandler

import (
	"api/internal/adapter/inbound/api/shared"
	"api/internal/core"
	"context"
	"net/http"
)

func ById(repo core.BitAlertRepo) http.Handler {
	getter := func(ctx context.Context, id string) (*core.Reaction, error) {
		return repo.GetBitAlertById(ctx, id)
	}
	setter := func(ctx context.Context, id string, r core.Reaction) error {
		return repo.SetBitAlertById(ctx, id, r)
	}
	deleter := func(ctx context.Context, id string) error { return repo.DeleteBitAlertById(ctx, id) }

	return shared.ReactionHandler{
		Getter: getter,
		Setter: setter,
		Delete: deleter,
	}
}

func List(repo core.BitAlertRepo) http.Handler {
	list := func(ctx context.Context) (map[string]core.Reaction, error) {
		return repo.GetAllBitAlerts(ctx)
	}
	setter := func(ctx context.Context, id string, r core.Reaction) error {
		return repo.SetBitAlertById(ctx, id, r)
	}

	return shared.ReactionList{
		List:   list,
		Setter: setter,
		Step:   100,
	}
}
