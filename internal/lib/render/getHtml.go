package render

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinrai-js/go/internal/lib/components"
	"github.com/jinrai-js/go/internal/lib/config"
	"github.com/jinrai-js/go/internal/lib/jinrai_value"
	"github.com/jinrai-js/go/internal/lib/pass"
	"github.com/jinrai-js/go/internal/lib/path_resolver"
	"github.com/jinrai-js/go/internal/tools"
)

func GetHTML(ctx context.Context, content *[]config.Content, keys []string) string {
	var result strings.Builder

	for _, chunk := range *content {
		str := renderChunk(ctx, &chunk, keys)
		result.WriteString(str)
	}

	return result.String()
}

func renderChunk(ctx context.Context, chunk *config.Content, keys []string) string {
	defer pass.Catch()

	switch chunk.Type {
	case "html":
		return tools.GetTemplate(ctx, chunk.TemplateName)

	case "value":
		value := path_resolver.GetValueByPath(ctx, chunk.Key, keys)
		str := fmt.Sprint(value)
		return str

	case "array":
		list := mapByKeys(ctx, func(key string) string {
			return GetHTML(ctx, &chunk.Data, append(keys, key))
		}, chunk.Key, keys)

		return strings.Join(list, "")

	case "custom":
		componentProps := jinrai_value.Parse(ctx, chunk.Props, keys)
		return components.Get(chunk.Name, componentProps)

	}

	return ""
}
