package jsonConfig

import (
	"path/filepath"
	"strings"

	"github.com/jinrai-js/go/internal/tools"
)

func (c Config) RenderIndex(html string, head string) []byte {
	index := tools.ReadHTML(filepath.Join(c.OutDir, "../index.html"))

	index = strings.Replace(index, "<!--app-html-->", html, 1)
	index = strings.Replace(index, "<!--app-head-->", head, 1)
	index = strings.ReplaceAll(index, "<!--->", "")

	return []byte(index)
}
