package jinrai

import (
	"fmt"
	"net/http"
)

func (c Static) Handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			// w.Write(templates.RenderIndex("panic", ""))
			fmt.Fprintf(w, "паника: %v", r)
		}
	}()

	c.Log("url: ", r.URL.Path)

	var route = c.Config.FindTemplateAndRender(r.URL)
	if route == nil {
		w.Write(c.Config.RenderIndex("", ""))
		c.Log("route nil")
		return
	}

	html, head := c.Generate(r.URL, route)

	w.Write(c.Config.RenderIndex(html, head))
}
