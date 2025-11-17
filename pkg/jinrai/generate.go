package jinrai

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/jinrai-js/go/pkg/jinrai/context"
	"github.com/jinrai-js/go/pkg/jinrai/jsonConfig"
)

func (c Static) Generate(url *url.URL, route *jsonConfig.Route) (string, string) {
	var context = context.New(url, c.OutDir)

	context.ExecuteRequests(c.Proxy, route.Requests, c.Rewrite)
	html := context.GetHTML(route.Content, []string{})

	return html, wrapExport(context.Output.Export)
}

func wrapExport(input any) string {
	export, _ := json.Marshal(input)
	result := fmt.Sprintf(`<script>window.next_f = %s</script>`, string(export))

	return result
}
