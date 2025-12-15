package render

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinrai-js/go/internal/lib/components"
	"github.com/jinrai-js/go/internal/lib/config"
	"github.com/jinrai-js/go/internal/lib/jinrai_value"
	"github.com/jinrai-js/go/internal/lib/path_resolver"
	"github.com/jinrai-js/go/internal/tools"
)

func GetHTML(ctx context.Context, content *config.Content, keys []string) string {
	var result strings.Builder

	for _, props := range *content {
		switch props.Type {
		case "html":
			result.WriteString(tools.GetTemplate(ctx, props.TemplateName))

		case "value":
			value := path_resolver.GetValueByPath(ctx, props.Key, keys)
			str := fmt.Sprint(value)
			result.WriteString(str)

		case "array":
			list := mapByKeys(ctx, func(key string) string {
				return GetHTML(ctx, &props.Data, append(keys, key))
			}, props.Key, keys)

			result.WriteString(strings.Join(list, ""))

		case "custom":
			componentProps := jinrai_value.Parse(ctx, props.Props, keys)
			result.WriteString(components.Get(props.Name, componentProps))

		}

	}

	return result.String()
}
