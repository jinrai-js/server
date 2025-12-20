package cashe

import "github.com/jinrai-js/go/internal/lru"

var data = lru.New(1000)

func Get(key string) (any, bool) {
	if val, exists := data.Get(key); exists == nil {
		return val, true
	}
	return nil, false
}

func Set(key string, value any) {
	data.Put(key, value)
}
