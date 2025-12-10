package jinrai

import (
	"fmt"
	"net/http"

	"github.com/jinrai-js/go/pkg/lib/handler"
	"github.com/jinrai-js/go/pkg/lib/requestScope"
	"github.com/jinrai-js/go/pkg/lib/serverState"
)

func (c *Jinrai) Handler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			// w.Write(templates.RenderIndex("panic", ""))
			fmt.Fprintf(w, "паника: %v", r)
		}
	}()

	c.Log("html url: ", r.URL.Path)

	var route = handler.FindTemplate(r.URL, &c.Json.Routes)

	if route == nil {
		// w.Write(c.Config.RenderIndex("", ""))
		c.Log("route nil")
		return
	}

	ctx := r.Context()
	ctx = c.With(ctx)
	ctx = requestScope.With(ctx, requestScope.New(r.URL.Path, r.URL.Query()))
	ctx = serverState.With(ctx, serverState.New(*c.Server.Proxy, route.State))

	handler.Render(ctx, route.Content)
	// html := render.GetHTML(ctx, route.Content, []string{})

	// log.Println(html)

	// w.Write(c.Config.RenderIndex(html, head))
}
