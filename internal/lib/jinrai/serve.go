package jinrai

import (
	"fmt"
	"net/http"
	"path/filepath"
	"slices"
	"strings"

	"github.com/jinrai-js/server/internal/extensions"
	proxymap "github.com/jinrai-js/server/internal/lib/proxy"
	"github.com/jinrai-js/server/internal/proxy"
)

func (c *Jinrai) Serve(port int) error {
	mux := http.NewServeMux()

	assets := http.FileServer(http.Dir(c.Server.Dist))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		if c.Server.Rewrite != nil {
			url = (*c.Server.Rewrite)(r.URL.Path)
		}
		url = strings.TrimRight(url, "/")

		if url == "" {
			url = "/"
		}

		if url == "/" {
			c.Handler(w, r, url)
			return
		}

		if c.Server.Proxy != nil {
			if target, prefix, find := proxymap.Find(url, c.Server.Proxy); find {
				proxy.Handler(w, r, prefix, target, c.Server.Verbose)
				return
			}
		}

		ext := filepath.Ext(url)
		if c.Server.Assets != nil && strings.Contains(url, ".") && slices.Contains(extensions.Popular, ext) {
			assets.ServeHTTP(w, r)
			return
		}

		c.Handler(w, r, url)
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
