package server_context

import (
	"context"

	"github.com/jinrai-js/go/pkg/lib/server_state"
)

type stateKey struct{}

func With(ctx context.Context, state server_state.State) context.Context {
	return context.WithValue(ctx, stateKey{}, &state)
}

func Get(ctx context.Context) *server_state.State {
	return ctx.Value(stateKey{}).(*server_state.State)
}
