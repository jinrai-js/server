package cashe

import (
	"sync"
)

var (
	cache = make(map[string]string)
	mutex = &sync.RWMutex{}
)

func GetValue(key string) (string, bool) {
	mutex.RLock()
	defer mutex.RUnlock()

	val, ok := cache[key]
	return val, ok
}

func SetValue(key string, value string) {
	mutex.Lock()
	defer mutex.Unlock()

	cache[key] = value
}
