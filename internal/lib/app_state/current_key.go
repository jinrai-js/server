package app_state

import (
	"context"
	"strings"

	"github.com/jinrai-js/server/internal/lib/jinrai_value"
)

func (s *AppState) GetCurrentKey(ctx context.Context, keys []string) string {
	return convertKeyToString(ctx, s.Key, keys)
}

func convertKeyToString(ctx context.Context, key any, keys []string) string {
	switch v := key.(type) {
	case string:
		return v

	case []string:
		return strings.Join(v, "-")

	case []any:
		result := []string{}
		for _, val := range jinrai_value.Parse(ctx, v, keys).([]any) {
			result = append(result, convertKeyToString(ctx, val, keys))
		}

		return convertKeyToString(ctx, result, keys)
	}

	return ""
}
