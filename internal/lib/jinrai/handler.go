package jinrai

import (
	"context"
	"net/http"

	"github.com/jinrai-js/server/internal/lib/config/app_context"
	"github.com/jinrai-js/server/internal/lib/handler"
	"github.com/jinrai-js/server/internal/lib/index"
	"github.com/jinrai-js/server/internal/lib/interfaces"
	"github.com/jinrai-js/server/internal/lib/jlog"
	"github.com/jinrai-js/server/internal/lib/meta"
	"github.com/jinrai-js/server/internal/lib/request"
	"github.com/jinrai-js/server/internal/lib/request/request_context"
	"github.com/jinrai-js/server/internal/lib/server_error"
	"github.com/jinrai-js/server/internal/lib/server_state"
	"github.com/jinrai-js/server/internal/lib/server_state/server_context"
)

func (c *Jinrai) Handler(w http.ResponseWriter, r *http.Request) {
	defer server_error.Catch()
	jlog.Writeln("👉 ", r.URL.Path)

	content, states := handler.FindTemplate(r.URL, &c.Json.Routes)
	ctx := c.CreateContext(r, states)

	meta.Load(ctx)

	if content == nil {
		w.Write(index.RenderIndex(ctx, c.Server.Dist, ""))
		return
	}

	html := handler.Render(ctx, content)

	w.Write(index.RenderIndex(ctx, c.Server.Dist, html))
}

func (c *Jinrai) CreateContext(r *http.Request, states interfaces.States) context.Context {
	ctx := r.Context()
	ctx = app_context.WithJson(ctx, &c.Json)
	ctx = app_context.WithServer(ctx, &c.Server)
	ctx = request_context.With(ctx, request.New(r.URL.Path, r.URL.Query(), r.URL.RawQuery))
	ctx = server_context.With(ctx, server_state.New(*c.Server.Proxy, states))

	return ctx
}
