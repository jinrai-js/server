package index

import (
	"context"
	"regexp"
	"strings"
	"sync"

	"github.com/jinrai-js/server/internal/lib/lang/lang_context"
	"github.com/jinrai-js/server/internal/tools"
)

type chunk struct {
	Type    string
	Value   string
	Default string
	Key     string
}

var (
	mu     sync.RWMutex
	isInit = false
	chunks []chunk
)

func getIndexChunks(file string) *[]chunk {
	mu.RLock()

	if isInit {
		defer mu.RUnlock()
		return &chunks
	}

	mu.RUnlock()
	mu.Lock()
	defer mu.Unlock()

	re := regexp.MustCompile(`(\{\{|\}\})`)

	index := tools.ReadHTML(file, false)

	for index, val := range re.Split(index, -1) {
		if index%2 == 0 {
			chunks = append(chunks, chunk{
				Type:  "HTML",
				Value: val,
			})
		} else {
			values := strings.SplitN(val, "|", 2)

			var def string
			if len(values) > 1 {
				def = values[1]
			}

			chunks = append(chunks, chunk{
				Type:    "META",
				Key:     values[0],
				Default: def,
			})
		}
	}

	isInit = true

	return &chunks
}

func (c *chunk) ToString(ctx context.Context, meta *map[string]string) string {
	switch c.Type {
	case "HTML":
		return c.Value

	case "META":
		return getMetaValue(ctx, meta, c.Key, c.Default)
	}

	return ""
}

func getMetaValue(ctx context.Context, meta *map[string]string, key, def string) string {
	if key == "lang" {
		lang := lang_context.Get(ctx)
		return lang.Active
	}

	if val, ok := (*meta)[key]; ok {
		return val
	}

	return def
}
