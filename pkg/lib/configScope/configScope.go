package configScope

import (
	"context"

	"github.com/jinrai-js/go/pkg/lib/jinrai"
)

type jinraiKey struct{}

func With(ctx context.Context, config *jinrai.Jinrai) context.Context {
	return context.WithValue(ctx, jinraiKey{}, config)
}

func Get(ctx context.Context) *jinrai.Jinrai {
	return ctx.Value(jinraiKey{}).(*jinrai.Jinrai)
}
