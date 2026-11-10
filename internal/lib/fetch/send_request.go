package fetch

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	"github.com/jinrai-js/server/internal/lib/proxy"
	"github.com/jinrai-js/server/internal/lib/request_cashe"
)

func AsyncSendRequest(ctx context.Context, url string, method string, body any) string {
	key := getKey(ctx, url, method, body)
	if value, exists := cashe.Get(ctx, key); exists {
		return value
	}

	fetch_group.Run(key)
	go func() {
		defer fetch_group.Done(key)

		result, err := SendRequest(ctx, url, method, body)
		if err == nil {
			cashe.Set(ctx, key, result)
		} else {
			app_error.Create(ctx, err)
		}
	}()

	pass.Exit()
	return ""
}

func getKey(ctx context.Context, url, method string, body any) string {
	jsonBody, _ := json.Marshal(body)

	useGlobalCashe := request_cashe.CheckUrl(ctx, url)
	var prefix string

	if !useGlobalCashe {
		prefix = "~"
	}

	return prefix + url + "|" + method + "|" + string(jsonBody)
}

// SendRequest отправит запрос на сервер с учетом Proxy
func SendRequest(ctx context.Context, url string, method string, body any) (string, error) {
	proxyUrl := proxy.GetUrl(ctx, url)

	if !strings.HasPrefix(proxyUrl, "http") {
		return "", fmt.Errorf("PROXY ERROR: Запрос всегда должне быть http или https (%s)", proxyUrl)
	}

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
