package app_error

import (
	"context"

	"github.com/jinrai-js/server/internal/lib/app_error/error_context"
)

func Create(ctx context.Context, err error) {
	scope := error_context.Get(ctx)
	scope.Message = err.Error()
	scope.Exists = true
}

func Has(ctx context.Context) bool {
	scope := error_context.Get(ctx)
	return scope.Exists
}

func Get(ctx context.Context) error_context.Scope {
	scope := error_context.Get(ctx)
	return *scope
}
