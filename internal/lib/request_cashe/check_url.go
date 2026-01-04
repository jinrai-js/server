package request_cashe

import (
	"context"
	"strings"

	"github.com/jinrai-js/server/internal/lib/config/app_context"
)

func CheckUrl(ctx context.Context, url string) bool {
	config := app_context.GetJson(ctx)
	for _, path := range config.CacheablePaths {
		if strings.HasSuffix(path, "*") {
			if strings.HasPrefix(url, path[:len(path)-1]) {
				return true
			}
		} else {
			if url == path {
				return true
			}
		}
	}

	return false
}
