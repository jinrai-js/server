package cashe

import (
	"context"
	"strings"

	"github.com/jinrai-js/server/internal/lib/global_cashe"
	"github.com/jinrai-js/server/internal/lib/request_cashe/request_cashe_context"
)

func Get(ctx context.Context, key string) (string, bool) {
	if isLocal(key) {
		cashe := request_cashe_context.Get(ctx)
		return cashe.Get(key)
	} else {
		return global_cashe.Get(key)
	}
}

func Set(ctx context.Context, key string, value string) {
	if isLocal(key) {
		cashe := request_cashe_context.Get(ctx)
		cashe.Set(key, value)
	} else {
		global_cashe.Set(key, value)
	}
}

func isLocal(key string) bool {
	return strings.HasPrefix(key, "~")
}
