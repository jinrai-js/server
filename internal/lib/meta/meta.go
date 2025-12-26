package meta

import (
	"context"
	"encoding/json"

	"github.com/jinrai-js/server/internal/lib/config/app_context"
	"github.com/jinrai-js/server/internal/lib/fetch"
	"github.com/jinrai-js/server/internal/lib/pass"
	"github.com/jinrai-js/server/internal/lib/request/request_context"
)

type MetaResponce struct {
	Data map[string]string `json:"data"`
}

func Load(ctx context.Context) string {
	defer pass.Catch()

	app := app_context.GetServer(ctx)
	request := request_context.Get(ctx)

	if app.Meta == nil {
		return ""
	}

	return fetch.AsyncSendRequest(ctx, *app.Meta, "POST", map[string]string{
		"route": request.Url,
	})

}

func Get(ctx context.Context) *map[string]string {
	app := app_context.GetServer(ctx)
	if app.Meta == nil {
		return nil
	}

	data := Load(ctx)

	var tags MetaResponce
	json.Unmarshal([]byte(data), &tags)

	return &tags.Data
}
