package custom_components

import (
	"context"

	"github.com/jinrai-js/go/pkg/lib/config/app_context"
)

func Render(ctx context.Context, name string, props any) string {
	server := app_context.GetServer(ctx)

	if handler, exists := (*server.Components)[name]; exists {
		return handler(props)
	}

	return "---"
}
