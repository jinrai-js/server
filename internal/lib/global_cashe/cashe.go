package global_cashe

import (
	"sync"

	"github.com/jinrai-js/server/internal/lru"
)

var (
	mu   sync.RWMutex
	data = lru.New(3000)
)

func Get(key string) (string, bool) {
	mu.RLock()
	defer mu.RUnlock()

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
