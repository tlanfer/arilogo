package core

import "context"

type GlobalConfigRepo interface {
	GetIdleState(ctx context.Context) (*State, error)
	SetIdleState(ctx context.Context, state State) error
}
