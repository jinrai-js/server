package index

import (
	"regexp"
	"strings"
	"sync"

	"github.com/jinrai-js/server/internal/tools"
)

type chunk struct {
	Type    string
	Value   string
	Default string
	Key     string
}

var (
	mu     sync.Mutex
	isInit = false
	chunks []chunk
)

func getIndexChunks(file string) *[]chunk {
	mu.Lock()
	defer mu.Unlock()

	if isInit {
		return &chunks
	}

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

func (c *chunk) ToString(meta *map[string]string) string {
	switch c.Type {
	case "HTML":
		return c.Value

	case "META":
		if val, ok := (*meta)[c.Key]; ok {
			return val
		}
		return c.Default
	}

	return ""
}
