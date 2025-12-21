package cashe

import (
	"sync"

	"github.com/jinrai-js/go/internal/lru"
)

var mu sync.Mutex
var data = lru.New(1000)

func Get(key string) (any, bool) {
	mu.Lock()
	defer mu.Unlock()

	if val, exists := data.Get(key); exists == nil {
		return val, true
	}
	return nil, false
}

func Set(key string, value any) {
	mu.Lock()
	defer mu.Unlock()

	data.Put(key, value)
}
