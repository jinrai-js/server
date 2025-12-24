package proxy

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/jinrai-js/server/internal/cashe"
)

// #FIX - не используется prefix - обернуть в поиск префикса
func Handler(w http.ResponseWriter, r *http.Request, prefix, targetURL string, verbose bool) {
	base, err := url.Parse(targetURL)
	if err != nil {
		http.Error(w, "Invalid target URL", http.StatusInternalServerError)
		return
	}

	target := base.ResolveReference(&url.URL{
		Path:     r.URL.Path,
		RawQuery: r.URL.RawQuery,
	})

	proxyReq, err := http.NewRequest(r.Method, target.String(), r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	proxyReq.Header = r.Header.Clone()

	key := cashe.GetRequestKey(proxyReq)

	if body, ok := cashe.GetValue(key); ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
		return
	}

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Proxy error: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for k, values := range resp.Header {
		for _, v := range values {
			w.Header().Add(k, v)
		}
	}
	w.WriteHeader(resp.StatusCode)

	if strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Read error", http.StatusInternalServerError)
			return
		}

		cashe.SetValue(key, string(respBody))
		w.Write(respBody)
		return
	}

	io.Copy(w, resp.Body)
}
