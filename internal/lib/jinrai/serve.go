package jinrai

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/jinrai-js/server/internal/lib/jlog"
	"github.com/jinrai-js/server/internal/proxy"
)

func (c *Jinrai) Serve(port int) error {
	mux := http.NewServeMux()

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

		if c.Server.Assets != nil {
			filePath := path.Join(c.Server.Dist, r.URL.Path)

			if _, err := os.Stat(filePath); err == nil {
				jlog.Writeln("📁 ", r.URL.Path)

				fs := http.FileServer(http.Dir(c.Server.Dist))
				fs.ServeHTTP(w, r)
				return
			}
		}

		c.Handler(w, r)
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
