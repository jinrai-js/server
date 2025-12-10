package render

import (
	"context"
	"strings"

	"github.com/jinrai-js/go/internal/tools"
	"github.com/jinrai-js/go/pkg/lib/server_state/server_context"
)

func GetValueByPath(ctx context.Context, path string, keys []string) any { // #TODO получить данные из server STATE
	split := strings.SplitN(path, "@", 2)
	stateKey := split[0]
	pathItems := strings.Split(split[1], "/")

	state := server_context.Get(ctx)
	link, exists := state.Get(stateKey, keys)
	if !exists {
		return nil
	}

	for index, pathItem := range pathItems {
		if index == 0 {
			pathItem = "data"
		}

		if strings.HasPrefix(pathItem, "[ITEM=") {
			if len(keys) == 0 {
				pathItem = pathItem[6 : len(pathItem)-1]
			} else {
				pathItem = keys[0]
				keys = keys[1:]
			}

		}

		switch v := link.(type) {
		case map[string]interface{}:
			if val, ok := v[pathItem]; ok {
				link = val
			} else {
				return ""
			}
		case map[int]interface{}:
			if val, ok := v[tools.StrToInt(pathItem)]; ok {
				link = val
			} else {
				return ""
			}
		case []interface{}:
			index := tools.StrToInt(pathItem)
			if index >= 0 && index < len(v) {
				link = v[index]
			}
		default:
			return ""
		}
	}

	return link
}

func getSliceByPath(ctx context.Context, path string, keys []string) []any {
	value := GetValueByPath(ctx, path, keys)
	if val, ok := value.([]any); ok {
		return val
	}

	return []any{}
}

func getStateValue(ctx context.Context, stateKey string) {

}
