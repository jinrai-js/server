package jinrai

import (
	"context"
)

type jinraiKey struct{}

func (c *Jinrai) With(ctx context.Context) context.Context {
	return context.WithValue(ctx, jinraiKey{}, c)
}

func Get(ctx context.Context) *Jinrai {
	if j, ok := ctx.Value(jinraiKey{}).(*Jinrai); ok {
		return j
	}
	return nil
}
