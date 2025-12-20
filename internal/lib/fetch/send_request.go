package fetch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jinrai-js/go/internal/lib/cashe"
	"github.com/jinrai-js/go/internal/lib/fetch_group"
	"github.com/jinrai-js/go/internal/lib/pass"
	"github.com/jinrai-js/go/internal/lib/server_state/server_context"
	"github.com/jinrai-js/go/internal/tools"
)

func AsyncSendRequest(ctx context.Context, url string, method string, body any) any {
	key := getKey(url, method, body)
	if value, exists := cashe.Get(key); exists {
		return value
	}

	fetch_group.Run(key)
	go func() {
		defer fetch_group.Done(key)

		result, err := SendRequest(ctx, url, method, body)
		if err == nil {
			cashe.Set(key, result)
		} else {
			log.Fatal("URL:", url, method, ">>", err)
		}
	}()

	pass.Exit()
	return nil
}

func getKey(url, method string, body any) string {
	jsonBody, _ := json.Marshal(body)

	return url + "|" + method + "|" + string(jsonBody)
}

// SendRequest отправит запрос на сервер с учетом Proxy + сначала выполнится проверка на cashe
func SendRequest(ctx context.Context, url string, method string, body any) (any, error) {
	proxyUrl := getUrlWithProxy(ctx, url)

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	log.Println(proxyUrl, string(jsonBody))

	req, err := http.NewRequest(method, proxyUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		// Timeout: time.Millisecond * 1000,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fmt.Println(" OK")

	result := tools.IoToJson(resp.Body)

	return result, nil
}

func getUrlWithProxy(ctx context.Context, url string) string {
	server := server_context.Get(ctx)

	for prefix, targetURL := range server.Proxy {
		if strings.HasPrefix(url, prefix) {
			result := strings.Replace(url, prefix, targetURL+prefix, 1)
			return result
		} else {
			// #TODO Что если Proxy нет?
		}
	}

	return url
}
