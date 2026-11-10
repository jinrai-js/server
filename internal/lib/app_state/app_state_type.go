package app_state

import (
	"context"
	"encoding/json"

	"github.com/jinrai-js/server/internal/lib/fetch"
	"github.com/jinrai-js/server/internal/lib/fetch_group"
	"github.com/jinrai-js/server/internal/lib/interfaces"
	"github.com/jinrai-js/server/internal/lib/jinrai_value"
	"github.com/jinrai-js/server/internal/lib/pass"
)

type StateRequest struct {
	Input  map[string]any `json:"input"`
	Method string         `json:"method"`
	Url    string         `json:"url"`
}

type StateSource struct {
	Request *StateRequest `json:"request,omitempty"`
}

type StateOption struct {
	Source   *StateSource `json:"source,omitempty"`
	Required bool         `json:"required"`
}

type AppState struct {
	Options StateOption `json:"options"`
	Key     any         `json:"key"`
	Value   any         `json:"value"`
}

//

func New(data map[string]any) states {
	var result states

	if str, err := json.Marshal(data); err == nil {
		json.Unmarshal(str, &result)
	}

	return result
}

type states map[string]AppState

// Get - получить State по названию
func (s *states) Get(name string) interfaces.State {
	if state, exist := (*s)[name]; exist {
		return &state
	}

	return nil
}

func (s *states) GetWithoutSource(ctx context.Context) map[string]any {
	result := make(map[string]any)

	s.InitRequired(ctx)
	fetch_group.Wait()

	for key, state := range *s {
		source := state.Options.Source
		if source == nil || state.Value != nil {
			result[key] = state.Value
		}

	}

	return result
}

func (s *states) InitRequired(ctx context.Context) {
	states := *s
	for key, state := range states {
		if !state.Options.Required {
			continue
		}

		initRequiredState(ctx, &state)
		states[key] = state
	}
}

func initRequiredState(ctx context.Context, state *AppState) {
	defer pass.Catch()

	request := state.Options.Source.Request

	currentInput := jinrai_value.Parse(ctx, request.Input, []string{})
	if v := fetch.AsyncSendRequest(ctx, request.Url, request.Method, currentInput); v != "" {
		var data any
		json.Unmarshal([]byte(v), &data)

		state.Value = data
	}
}
