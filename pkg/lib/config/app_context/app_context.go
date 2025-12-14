package app_context

import (
	"context"

	"github.com/jinrai-js/go/pkg/lib/config"
)

type jsonKey struct{}

func WithJson(ctx context.Context, json *config.JsonConfig) context.Context {
	return context.WithValue(ctx, jsonKey{}, json)
}

func GetJson(ctx context.Context) *config.JsonConfig {
	if json, ok := ctx.Value(jsonKey{}).(*config.JsonConfig); ok {
		return json
	}
	return nil
}

//// server

type serverKey struct{}

func WithServer(ctx context.Context, server *config.Server) context.Context {
	return context.WithValue(ctx, serverKey{}, server)
}

func GetServer(ctx context.Context) *config.Server {
	if json, ok := ctx.Value(serverKey{}).(*config.Server); ok {
		return json
	}
	return nil
}
