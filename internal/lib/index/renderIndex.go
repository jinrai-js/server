package index

import (
	"path/filepath"
	"strings"

	"github.com/jinrai-js/server/internal/tools"
)

func RenderIndex(dist string, html string, head string) []byte {

	index := tools.ReadHTML(filepath.Join(dist, "index.html"))

	index = strings.Replace(index, "<!--app-html-->", html, 1)
	index = strings.Replace(index, "<!--app-head-->", head, 1)
	index = strings.ReplaceAll(index, "<!--->", "")

	return []byte(index)
}
