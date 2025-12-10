package render

import (
	"context"
	"strconv"
)

func mapByKeys(ctx context.Context, callback func(key string) string, path string, keys []string) []string {
	var result []string

	for key := range getSliceByPath(ctx, path, keys) {
		result = append(result, callback(strconv.Itoa(key)))
	}

	return result
}
