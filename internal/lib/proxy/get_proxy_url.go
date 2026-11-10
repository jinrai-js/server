package proxy

import (
	"context"
	"strings"

	"github.com/jinrai-js/server/internal/lib/server_state/server_context"
)

func GetUrl(ctx context.Context, url string) string {
	server := server_context.Get(ctx)

	target, _, _ := Find(url, &server.Proxy)
	return target
}

func Find(url string, proxy *map[string]string) (string, string, bool) {
	for prefix, targetURL := range *proxy {
		if strings.HasPrefix(url, prefix) {
			repl := targetURL
			if !strings.HasSuffix(prefix, "/") {
				repl += prefix
			}
			result := strings.Replace(url, prefix, repl, 1)
			return result, prefix, true
		} else {
			// #TODO Что если Proxy нет?
		}
	}

	return url, "", false
}
