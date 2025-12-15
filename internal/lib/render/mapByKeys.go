package render

import (
	"context"
	"strconv"

	"github.com/jinrai-js/go/internal/lib/path_resolver"
)

func mapByKeys(ctx context.Context, callback func(key string) string, path string, keys []string) []string {
	var result []string

	for key := range path_resolver.GetSliceByPath(ctx, path, keys) {
		result = append(result, callback(strconv.Itoa(key)))
	}

	return result
}
