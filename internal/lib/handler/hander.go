package handler

import (
	"context"
	"net/url"

	"github.com/jinrai-js/go/internal/lib/config"
	"github.com/jinrai-js/go/internal/lib/interfaces"
	"github.com/jinrai-js/go/internal/lib/render"
)

func Render(ctx context.Context, content *config.Content) string {
	html := render.GetHTML(ctx, content, []string{})

	return html
}

func FindTemplate(url *url.URL, routes *[]config.Route) (*config.Content, interfaces.States) {
	return render.FindTemplateAndRender(url, routes)
}

func GetValueByPath(ctx context.Context, path string, keys []string) any {
	return render.GetValueByPath(ctx, path, keys)
}
