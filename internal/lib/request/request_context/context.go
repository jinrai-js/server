package request_context

import (
	"context"

	"github.com/jinrai-js/go/internal/lib/request"
)

type scopeKey struct{}

func With(ctx context.Context, scope request.Scope) context.Context {
	return context.WithValue(ctx, scopeKey{}, &scope)
}

func Get(ctx context.Context) *request.Scope {
	return ctx.Value(scopeKey{}).(*request.Scope)
}
