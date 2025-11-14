package server

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

func Run(port int, handler func(w http.ResponseWriter, r *http.Request), assetsPath *string, proxyMap *map[string]string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if proxyMap != nil {
			for prefix, targetURL := range *proxyMap {
				if strings.HasPrefix(r.URL.Path, prefix) {
					proxyHandler(w, r, prefix, targetURL)
					return
				}
			}
		}

		if assetsPath != nil {
			filePath := path.Join(*assetsPath, r.URL.Path)

			if r.URL.Path == "/" {
				handler(w, r)
				return
			}

			if _, err := os.Stat(filePath); err == nil {
				fs := http.FileServer(http.Dir(*assetsPath))
				fs.ServeHTTP(w, r)
				return
			}
		}

		handler(w, r)
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func proxyHandler(w http.ResponseWriter, r *http.Request, prefix string, targetURL string) {

	target, err := url.Parse(targetURL)
	if err != nil {
		http.Error(w, "Invalid target URL", http.StatusInternalServerError)
		return
	}

	proxyReq := &http.Request{
		Method: r.Method,
		URL:    target,
		Header: r.Header.Clone(),
		Body:   r.Body,
	}

	pathWithoutPrefix := strings.TrimPrefix(r.URL.Path, prefix)
	proxyReq.URL.Path = path.Join(target.Path, pathWithoutPrefix)

	proxyReq.URL.RawQuery = r.URL.RawQuery

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Proxy error: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}
