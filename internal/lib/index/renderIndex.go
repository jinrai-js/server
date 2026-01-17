package index

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/jinrai-js/server/internal/lib/meta"
	"github.com/jinrai-js/server/internal/lib/server_state/server_context"
)

func RenderIndex(ctx context.Context, dist string, html string) []byte {
	metatags := meta.Get(ctx)

	index := getIndex(ctx, filepath.Join(dist, "index.html"), metatags)

	server := server_context.Get(ctx)
	script := server.ExportScript(ctx)

	index = strings.Replace(index, "<!--app-html-->", html, 1)
	index = strings.Replace(index, "<!--app-head-->", script, 1)

	index = strings.ReplaceAll(index, "<!--dev-only", "")
	index = strings.ReplaceAll(index, "dev-only-->", "")

	return []byte(index)
}

func getIndex(ctx context.Context, file string, meta *map[string]string) string {
	chunks := getIndexChunks(file)

	var result strings.Builder

	for _, chunk := range *chunks {
		result.WriteString(chunk.ToString(ctx, meta))
	}

	return result.String()
}
