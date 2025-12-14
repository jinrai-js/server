package app_state

import (
	"context"
	"strings"

	"github.com/jinrai-js/go/pkg/lib/jinrai_value"
)

func (s *AppState) GetCurrentKey(ctx context.Context, keys []string) string {
	return convertKeyToString(ctx, s.Key)
}

func convertKeyToString(ctx context.Context, key any) string {
	switch v := key.(type) {
	case string:
		return v

	case []string:
		return strings.Join(v, "-")

	case []any:
		result := []string{}
		for _, val := range jinrai_value.Parse(ctx, v).([]any) {
			result = append(result, convertKeyToString(ctx, val))
		}

		return convertKeyToString(ctx, result)
	}

	return ""
}
