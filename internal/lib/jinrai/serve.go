package jinrai

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/jinrai-js/server/internal/proxy"
)

func (c *Jinrai) Serve(port int) error {
	mux := http.NewServeMux()

	assets := http.FileServer(http.Dir(c.Server.Dist))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			c.Handler(w, r)
			return
		}

		if c.Server.Proxy != nil {
			for prefix, targetURL := range *c.Server.Proxy {
				if strings.HasPrefix(r.URL.Path, prefix) {
					proxy.Handler(w, r, prefix, targetURL, c.Server.Verbose)
					return
				}
			}
		}

		if c.Server.Assets != nil && strings.Contains(r.URL.Path, ".") && len(filepath.Ext(r.URL.Path)) < 5 {
			assets.ServeHTTP(w, r)
			return
		}

		c.Handler(w, r)
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
