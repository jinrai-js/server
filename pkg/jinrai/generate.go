package jinrai

import (
	"net/url"

	"github.com/jinrai-js/go/pkg/jinrai/context"
	"github.com/jinrai-js/go/pkg/jinrai/jsonConfig"
)

func (c Static) Generate(url *url.URL, route *jsonConfig.Route) (body, head string) {
	var context = context.New(url, c.OutDir)

	body = context.GetHTML(route.Content, []string{})
	head = context.ServerState.Export()

	return
}
