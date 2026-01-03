package lang_base

import (
	"context"
	"encoding/json"
	"strings"
	"sync"

	"golang.org/x/sync/singleflight"

	"github.com/jinrai-js/server/internal/lib/fetch"
	"github.com/jinrai-js/server/internal/lib/jlog"
)

type langs map[string]map[string]string

var (
	cashe sync.Map
	group singleflight.Group
)

func GetValue(ctx context.Context, source, lang, key string) string {
	if v, ok := cashe.Load(lang); ok {
		return readMap(v, key)
	}

	v, err, _ := group.Do(lang, func() (any, error) {
		if v, ok := cashe.Load(lang); ok {
			return v, nil
		}

		url := strings.Replace(source, "*", lang, 1)
		body, err := fetch.SendRequest(ctx, url, "GET", nil)
		if err != nil {
			return map[string]string{}, err
		}

		var dict map[string]string
		if err := json.Unmarshal([]byte(body), &dict); err != nil {
			return map[string]string{}, err
		}

		cashe.Store(lang, dict)
		return dict, nil
	})

	if err != nil {
		jlog.Writeln("[LANG ERR]", err)
		return key
	}

	return readMap(v, key)
}

func readMap(v any, key string) string {
	dist := v.(map[string]string)
	if val, ok := dist[key]; ok {
		return val
	}
	return key
}
