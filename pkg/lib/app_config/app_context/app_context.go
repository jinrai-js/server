package app_context

import (
	"context"

	"github.com/jinrai-js/go/pkg/lib/app_config"
)

type jsonKey struct{}

func WithJson(ctx context.Context, json *app_config.JsonConfig) context.Context {
	return context.WithValue(ctx, jsonKey{}, json)
}

func GetJson(ctx context.Context) *app_config.JsonConfig {
	if json, ok := ctx.Value(jsonKey{}).(*app_config.JsonConfig); ok {
		return json
	}
	return nil
}

//// server

type serverKey struct{}

func WithServer(ctx context.Context, server *app_config.Server) context.Context {
	return context.WithValue(ctx, serverKey{}, server)
}

func GetServer(ctx context.Context) *app_config.Server {
	if json, ok := ctx.Value(serverKey{}).(*app_config.Server); ok {
		return json
	}
	return nil
}
