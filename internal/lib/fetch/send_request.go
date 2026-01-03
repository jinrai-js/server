package fetch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jinrai-js/server/internal/lib/app_error"
	"github.com/jinrai-js/server/internal/lib/cashe"
	"github.com/jinrai-js/server/internal/lib/fetch_group"
	"github.com/jinrai-js/server/internal/lib/jlog"
	"github.com/jinrai-js/server/internal/lib/pass"
	"github.com/jinrai-js/server/internal/lib/server_state/server_context"
)

func AsyncSendRequest(ctx context.Context, url string, method string, body any) string {
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
			app_error.Create(ctx, err)
		}
	}()

	pass.Exit()
	return ""
}

func getKey(url, method string, body any) string {
	jsonBody, _ := json.Marshal(body)

	return url + "|" + method + "|" + string(jsonBody)
}

// SendRequest отправит запрос на сервер с учетом Proxy + сначала выполнится проверка на cashe
func SendRequest(ctx context.Context, url string, method string, body any) (string, error) {
	proxyUrl := getUrlWithProxy(ctx, url)

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	jlog.Writeln("🤙 ", proxyUrl, string(jsonBody))

	req, err := http.NewRequest(method, proxyUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	client := &http.Client{
		Timeout: time.Millisecond * 3000,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		jlog.Writeln("❌")
		log.Panic(err)
	}

	if resp.Status != "200 OK" {
		return "", errors.New(string(bodyBytes))
	}

	return string(bodyBytes), nil
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
