package cashe

import (
	"sync"

	"github.com/jinrai-js/go/internal/lru"
)

var mu sync.Mutex
var data = lru.New(1000)

func Get(key string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()

	if val, exists := data.Get(key); exists == nil {
		return val, true
	}
	return "", false
}

func Set(key string, value string) {
	mu.Lock()
	defer mu.Unlock()

	data.Put(key, value)
}
