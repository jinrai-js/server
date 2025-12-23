package server_state

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jinrai-js/go/internal/lib/interfaces"
	"github.com/jinrai-js/go/internal/lib/server_error"
)

type State struct {
	State     map[string]any
	Proxy     map[string]string
	AppStates interfaces.States
}

func New(proxy map[string]string, state interfaces.States) State {
	return State{
		make(map[string]any),
		proxy,
		state,
	}
}

func (s *State) Get(ctx context.Context, stateName string, keys []string) (any, bool) {
	appState := s.AppStates.Get(stateName)

	if appState != nil {
		// currentKey ключ данных в state с учетом keys
		// (ключ ссылается на данные из конкретного запроса)
		currentKey := appState.GetCurrentKey(ctx, keys)

		if value, exists := s.State[currentKey]; exists {
			return value, true
		}

		if value, exists := appState.GetValue(ctx, keys); exists {
			s.State[currentKey] = value
			return value, true
		}
	}

	return nil, false
}

func (s *State) Export() string {
	export, _ := json.Marshal(map[string]any{
		"state":  s.State,
		"errors": server_error.Export(),
	})
	result := fmt.Sprintf(`<script>window.__appc__ = %s</script>`, string(export))

	return result
}
