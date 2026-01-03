package error_context

import (
	"context"
)

type scopeKey struct{}
type Scope struct {
	Message string `json:"message"`
	Exists  bool
}

func With(ctx context.Context) context.Context {
	scope := Scope{Exists: false}

	return context.WithValue(ctx, scopeKey{}, &scope)
}

func Get(ctx context.Context) *Scope {
	return ctx.Value(scopeKey{}).(*Scope)
}
