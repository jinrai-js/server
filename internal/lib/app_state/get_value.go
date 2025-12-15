package app_state

import (
	"context"

	"github.com/jinrai-js/go/internal/lib/fetch"
	"github.com/jinrai-js/go/internal/lib/jinrai_value"
)

// GetValue получить значение на основе StateInterface Option
// 1) отправляет запрос на сервер (sourse.request)
func (s *AppState) GetValue(ctx context.Context, keys []string) (any, bool) {
	// # TODO Добавить проверку на другие источники (но сейчас источник один)
	request, exists := s.GetSourceRequest()
	if !exists {
		return nil, false
	}

	currentInput := jinrai_value.Parse(ctx, request.Input, keys)
	if result, ok := fetch.SendRequest(ctx, request.Url, request.Method, currentInput); ok {
		return result, true
	}

	return nil, false
}

func (s *AppState) GetSourceRequest() (*StateRequest, bool) {
	req := s.Options.Source.Request
	if req == nil {
		return nil, false
	}

	return req, true
}
