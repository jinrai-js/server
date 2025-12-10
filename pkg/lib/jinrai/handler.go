package jinrai

import (
	"fmt"
	"net/http"

	"github.com/jinrai-js/go/pkg/lib/app_config/app_context"
	"github.com/jinrai-js/go/pkg/lib/handler"
	"github.com/jinrai-js/go/pkg/lib/request"
	"github.com/jinrai-js/go/pkg/lib/request/request_context"
	"github.com/jinrai-js/go/pkg/lib/server_state"
	"github.com/jinrai-js/go/pkg/lib/server_state/server_context"
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

	ctx = app_context.WithJson(ctx, &c.Json)
	ctx = app_context.WithServer(ctx, &c.Server)

	ctx = request_context.With(ctx, request.New(r.URL.Path, r.URL.Query()))
	ctx = server_context.With(ctx, server_state.New(*c.Server.Proxy, route.State))

	handler.Render(ctx, route.Content)
	// html := render.GetHTML(ctx, route.Content, []string{})

	// log.Println(html)

	// w.Write(c.Config.RenderIndex(html, head))
}
