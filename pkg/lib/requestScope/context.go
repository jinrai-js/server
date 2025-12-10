package requestScope

import "context"

type scopeKey struct{}

func With(ctx context.Context, scope Scope) context.Context {
	return context.WithValue(ctx, scopeKey{}, &scope)
}

func Get(ctx context.Context) *Scope {
	return ctx.Value(scopeKey{}).(*Scope)
}
