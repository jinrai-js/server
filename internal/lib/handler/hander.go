package handler

import (
	"context"
	"net/url"

	"github.com/jinrai-js/server/internal/lib/app_error"
	"github.com/jinrai-js/server/internal/lib/config"
	"github.com/jinrai-js/server/internal/lib/fetch_group"
	"github.com/jinrai-js/server/internal/lib/interfaces"
	"github.com/jinrai-js/server/internal/lib/path_resolver"
	"github.com/jinrai-js/server/internal/lib/render"
	"github.com/jinrai-js/server/internal/lib/server_error"
)

func Render(ctx context.Context, content *[]config.Content) string {
	defer server_error.Catch()

	level := 0
	for {
		html := render.GetHTML(ctx, content, []string{})

		fetch_group.Wait()

		if app_error.Has(ctx) {
			return ""
		}

		if !fetch_group.WasSandRequest() {
			return html
		} else {
			fetch_group.Reset()
		}

		level++
		if level > 3 {
			return ""
		}
	}
}

func FindTemplate(url *url.URL, routes *[]config.Route) (*[]config.Content, interfaces.States) {
	return render.FindTemplateAndRender(url, routes)
}

func GetValueByPath(ctx context.Context, path string, keys []string) any {
	return path_resolver.GetValueByPath(ctx, path, keys)
}
