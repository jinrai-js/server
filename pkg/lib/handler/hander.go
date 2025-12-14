package handler

import (
	"context"
	"net/url"

	"github.com/jinrai-js/go/pkg/lib/config"
	"github.com/jinrai-js/go/pkg/lib/render"
)

func Render(ctx context.Context, content config.Content) string {
	html := render.GetHTML(ctx, content, []string{})

	return html
}

func FindTemplate(url *url.URL, routes *[]config.Route) *config.App {
	return render.FindTemplateAndRender(url, routes)
}
