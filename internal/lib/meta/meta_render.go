package meta

import (
	"context"

	"github.com/jinrai-js/server/internal/lib/server_state/server_context"
)

func Render(ctx context.Context) string {
	metaTags := renderMetaDate(ctx)
	server := server_context.Get(ctx)

	return metaTags + "\n\n" + server.Export()
}
