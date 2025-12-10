package serverState

import "context"

type stateKey struct{}

func With(ctx context.Context, state State) context.Context {
	return context.WithValue(ctx, stateKey{}, &state)
}

func Get(ctx context.Context) *State {
	return ctx.Value(stateKey{}).(*State)
}
