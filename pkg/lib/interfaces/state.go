package interfaces

import "context"

type State interface {
	GetCurrentKey(ctx context.Context, keys []string) string
	GetValue(ctx context.Context) (any, bool)
}

type States interface {
	Get(name string) State
}
