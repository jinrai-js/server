package lang_context

import (
	"context"
)

type scopeKey struct{}

type Lang struct {
	Active    string
	SourceUrl string
	Default   string
}

func With(ctx context.Context, lang Lang) context.Context {
	return context.WithValue(ctx, scopeKey{}, &lang)
}

func Get(ctx context.Context) *Lang {
	return ctx.Value(scopeKey{}).(*Lang)
}
