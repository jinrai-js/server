package handler

import (
	"context"
	"net/url"

	"github.com/jinrai-js/go/pkg/lib/app_config"
	"github.com/jinrai-js/go/pkg/lib/render"
)

func Render(ctx context.Context, content app_config.Content) string {
	html := render.GetHTML(ctx, content, []string{})

	return html
}

func FindTemplate(url *url.URL, routes *[]app_config.Route) *app_config.Route {
	return render.FindTemplateAndRender(url, routes)
}
