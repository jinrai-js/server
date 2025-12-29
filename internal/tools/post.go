package tools

import (
	"bytes"
	"net/http"
	"time"

	"github.com/jinrai-js/server/internal/lib/jlog"
)

var cache = make(map[string]any)

func Post(url string, jsonBody string) (any, bool) {
	if val, ok := cache[url+"|"+jsonBody]; ok {
		return val, true
	}
	jlog.Writeln("[POST]", url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonBody)))
	if err != nil {
		// panic("Не удалось создать запрос " + url)
		return nil, false
	}

	client := &http.Client{
		Timeout: time.Millisecond * 1000,
	}
	resp, err := client.Do(req)
	if err != nil {
		// panic("Не удалось выполнить запрос " + err.Error())
		return nil, false
	}
	defer resp.Body.Close()

	jlog.Writeln("[POST]", "OK")

	result := IoToJson(resp.Body)
	cache[url+"|"+jsonBody] = result

	return result, true
}

type AsyncResult struct {
	Result any
	Url    string
	Ok     bool
}

func AsyncPost(url string, jsonBody string, ch chan AsyncResult) {
	result, ok := Post(url, jsonBody)

	ch <- AsyncResult{
		Result: result,
		Url:    url,
		Ok:     ok,
	}
}
