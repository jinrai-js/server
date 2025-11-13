package jinrai

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (c Static) Handler(w http.ResponseWriter, r *http.Request) {
	if exceptUrl(r.URL) {
		w.WriteHeader(400)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			// w.Write(templates.RenderIndex("panic", ""))
			fmt.Fprintf(w, "паника: %v", r)
		}
	}()

	var route = c.Config.FindTemplateAndRender(r.URL)
	if route == nil {
		w.Write(c.Config.RenderIndex("", ""))
		return
	}

	// if c.Verbose {
	log.Println("url:", r.URL.Query())
	// }

	html, head := c.Generate(r.URL, route)

	w.Write(c.Config.RenderIndex(html, head))
}

func exceptUrl(url *url.URL) bool {
	for _, format := range []string{"svg", "jpg", "png", "js", "css", "ico", "json"} {
		if strings.HasSuffix(url.Path, "."+format) {
			return true
		}
	}

	return false
}
