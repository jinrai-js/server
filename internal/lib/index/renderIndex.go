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

	index := getIndex(filepath.Join(dist, "index.html"), metatags)

	server := server_context.Get(ctx)
	script := server.ExportScript()

	index = strings.Replace(index, "<!--app-html-->", html, 1)
	index = strings.Replace(index, "<!--app-head-->", script, 1)
	index = strings.ReplaceAll(index, "<!--->", "")

	return []byte(index)
}

func getIndex(file string, meta *map[string]string) string {
	chunks := getIndexChunks(file)

	var result strings.Builder

	for _, chunk := range *chunks {
		result.WriteString(chunk.ToString(meta))
	}

	return result.String()
}
