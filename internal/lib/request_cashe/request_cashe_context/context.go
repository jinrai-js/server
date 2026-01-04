package request_cashe_context

import (
	"context"

	"github.com/jinrai-js/server/internal/lib/request_cashe"
)

type stateKey struct{}

func With(ctx context.Context, state request_cashe.Cashe) context.Context {
	return context.WithValue(ctx, stateKey{}, &state)
}

func Get(ctx context.Context) *request_cashe.Cashe {
	return ctx.Value(stateKey{}).(*request_cashe.Cashe)
}
