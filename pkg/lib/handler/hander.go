package handler

import (
	"context"
	"net/url"

	"github.com/jinrai-js/go/pkg/lib/appConfig"
	"github.com/jinrai-js/go/pkg/lib/render"
)

func Render(ctx context.Context, content appConfig.Content) string {
	html := render.GetHTML(ctx, content, []string{})

	return html
}

func FindTemplate(url *url.URL, routes *[]appConfig.Route) *appConfig.Route {
	return render.FindTemplateAndRender(url, routes)
}
