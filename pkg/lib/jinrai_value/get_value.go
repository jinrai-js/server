package jinrai_value

import (
	"context"
	"log"
	"strings"

	"github.com/jinrai-js/go/pkg/lib/request/request_context"
)

func (jv *JV) GetValue(ctx context.Context) any {
	switch jv.Type {
	case "searchArray":
		return jv.GetSearchArray(ctx)
	case "searchString":
		return jv.GetSearchString(ctx)
	case "proxy":
		return jv.GetProxy(ctx)
	case "searchFull":
		return jv.GetSearchFull(ctx)
	case "paramsIndex":
		return jv.GetParamsIndex(ctx)
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
func (jv *JV) GetProxy(ctx context.Context) string {
	return ""
}

// GetSearchFull получить все get параметры единой строкой
func (jv *JV) GetSearchFull(ctx context.Context) string {
	return ""
}

// GetParamsIndex получить позицию path в url
func (jv *JV) GetParamsIndex(ctx context.Context) string {
	return ""
}
