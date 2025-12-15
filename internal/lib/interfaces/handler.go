package interfaces

import "context"

type Handler interface {
	GetValueByPath(ctx context.Context, path string, keys []string) any
}
