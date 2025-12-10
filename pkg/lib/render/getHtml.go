package render

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinrai-js/go/internal/tools"
	"github.com/jinrai-js/go/pkg/lib/app_config"
)

func GetHTML(ctx context.Context, content app_config.Content, keys []string) string {
	var result strings.Builder

	for _, props := range content {
		switch props.Type {
		case "html":
			result.WriteString(tools.GetTemplate(ctx, props.TemplateName))

		case "value":
			value := GetValueByPath(ctx, props.Key, keys)
			str := fmt.Sprint(value)
			result.WriteString(str)

		case "array":
			list := mapByKeys(ctx, func(key string) string {
				return GetHTML(ctx, props.Data, append(keys, key))
			}, props.Key, keys)

			result.WriteString(strings.Join(list, ""))

		case "custom":
			result.WriteString("[custom]")

		}

	}

	return result.String()
}
