package jinrai_value

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/jinrai-js/server/internal/lib/path_resolver"
	"github.com/jinrai-js/server/internal/lib/request/request_context"
)

func (jv *JV) GetValue(ctx context.Context, keys []string) any {
	switch jv.Type {
	case "searchArray":
		return jv.GetSearchArray(ctx) // OK
	case "searchString":
		return jv.GetSearchString(ctx) // OK
	case "proxy":
		return jv.GetProxy(ctx, keys) // OK
	case "searchFull":
		return jv.GetSearchFull(ctx) // OK
	case "paramsIndex":
		return jv.GetParamsIndex(ctx) // OK
	}

	log.Panic("Неизвестный тип: jv." + jv.Type)
	return nil
}

// GetSearchString получить get параметр
func (jv *JV) GetSearchString(ctx context.Context) string {
	scope := request_context.Get(ctx)

	if !scope.Search.Has(jv.Key) {
		return jv.Def.(string)
	}

	return scope.Search.Get(jv.Key)
}

// GetSearchArray получить get параметр разбитый по разделителю
func (jv *JV) GetSearchArray(ctx context.Context) []string {
	str := jv.GetSearchString(ctx)
	return strings.Split(str, jv.Separator)
}

// GetProxy получить значение из ServerState
func (jv *JV) GetProxy(ctx context.Context, keys []string) any {
	value := path_resolver.GetValueByPath(ctx, jv.Key, keys)
	return value
}

// GetSearchFull получить все get параметры единой строкой
func (jv *JV) GetSearchFull(ctx context.Context) string {
	requests := request_context.Get(ctx)

	return requests.SearchFull
}

// GetParamsIndex получить позицию path в url
func (jv *JV) GetParamsIndex(ctx context.Context) string {
	requests := request_context.Get(ctx)

	index, err := strconv.Atoi(jv.Key)
	if err != nil {
		return ""
	}

	if index >= 0 && index < len(requests.Path) {
		return requests.Path[index]
	}

	return ""
}
