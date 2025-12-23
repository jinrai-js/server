package meta

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jinrai-js/go/internal/lib/config/app_context"
	"github.com/jinrai-js/go/internal/lib/fetch"
	"github.com/jinrai-js/go/internal/lib/pass"
	"github.com/jinrai-js/go/internal/lib/request/request_context"
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

func renderMetaDate(ctx context.Context) string {
	app := app_context.GetServer(ctx)
	if app.Meta == nil {
		return ""
	}

	data := Load(ctx)

	var tags MetaResponce
	json.Unmarshal([]byte(data), &tags)
	result := metaToStr(tags.Data)

	return result
}

func metaToStr(tags map[string]string) string {
	var result strings.Builder

	for name, value := range tags {
		var tag string
		if name == "title" {
			tag = fmt.Sprintf("<title>%s</title>", value)
		} else if strings.HasPrefix(name, "og:") {
			tag = fmt.Sprintf("<meta property=\"%s\" content=\"%s\">", name, value)
		} else {
			tag = fmt.Sprintf("<meta name=\"%s\" content=\"%s\">", name, value)
		}

		result.WriteString(tag)
		result.WriteString("\n")
	}

	return result.String()
}
