package request_cashe

import "sync"

type Cashe struct {
	mu   sync.RWMutex
	data map[string]string
}

func New() Cashe {
	return Cashe{
		data: make(map[string]string),
	}
}

func (c *Cashe) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

func (c *Cashe) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.data[key]
	return val, ok
}
