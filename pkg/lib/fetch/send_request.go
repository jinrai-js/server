package fetch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jinrai-js/go/internal/tools"
	"github.com/jinrai-js/go/pkg/lib/server_state/server_context"
)

// SendRequest отправит запрос на сервер с учетом Proxy + сначала выполнится проверка на cashe
func SendRequest(ctx context.Context, url string, method string, body any) (any, bool) {
	// # TODO добавить smart(не все запросы нужно) LRU cashe для кэширования запросов

	// if val, ok := cache[url+"|"+jsonBody]; ok {
	// 	return val, true
	// }
	// fmt.Print(">> " + url)

	proxyUrl := getUrlWithProxy(ctx, url)

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, false
	}

	req, err := http.NewRequest(method, proxyUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, false
	}

	client := &http.Client{
		Timeout: time.Millisecond * 1000,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, false
	}
	defer resp.Body.Close()

	fmt.Println(" OK")

	result := tools.IoToJson(resp.Body)
	// cache[url+"|"+jsonBody] = result

	return result, true
}

func getUrlWithProxy(ctx context.Context, url string) string {
	server := server_context.Get(ctx)

	for prefix, targetURL := range server.Proxy {
		if strings.HasPrefix(url, prefix) {
			// # TODO (ВАЖНО!) проверить замену url на proxy
			result := strings.Replace(url, prefix, targetURL, 1)
			return result
		}
	}

	return url
}
