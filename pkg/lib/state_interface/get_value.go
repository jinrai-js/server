package state_interface

import (
	"context"

	"github.com/jinrai-js/go/pkg/lib/fetch"
	"github.com/jinrai-js/go/pkg/lib/jinrai_value"
)

// GetValue получить значение на основе StateInterface Option
// 1) отправляет запрос на сервер (sourse.request)
func (s *StateInterface) GetValue(ctx context.Context) (any, bool) {
	// # TODO Добавить проверку на другие источники (но сейчас источник один)
	request, exists := s.GetSourceRequest()
	if !exists {
		return nil, false
	}

	currentInput := jinrai_value.Parse(ctx, request.Input)
	if result, ok := fetch.SendRequest(ctx, request.Url, request.Method, currentInput); ok {
		return result, true
	}

	return nil, false
}

func (s *StateInterface) GetSourceRequest() (*StateRequest, bool) {
	req := s.Options.Source.Request
	if req == nil {
		return nil, false
	}

	return req, true
}
