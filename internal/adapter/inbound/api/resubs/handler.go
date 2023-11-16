package resubsHandler

import (
	"api/internal/adapter/inbound/api/shared"
	"api/internal/core"
	"context"
	"net/http"
)

func ById(repo core.ReSubAlertRepo) http.Handler {
	getter := func(ctx context.Context, id string) (*core.Reaction, error) {
		return repo.GetResubById(ctx, id)
	}
	setter := func(ctx context.Context, id string, r core.Reaction) error {
		return repo.SetResubById(ctx, id, r)
	}
	deleter := func(ctx context.Context, id string) error { return repo.DeleteResubById(ctx, id) }

	return shared.ReactionHandler{
		Getter: getter,
		Setter: setter,
		Delete: deleter,
	}
}

func List(repo core.ReSubAlertRepo) http.Handler {
	list := func(ctx context.Context) (map[string]core.Reaction, error) {
		return repo.GetAllReSubs(ctx)
	}
	setter := func(ctx context.Context, id string, r core.Reaction) error {
		return repo.SetResubById(ctx, id, r)
	}

	return shared.ReactionList{
		List:   list,
		Setter: setter,
		Step:   1,
	}
}
