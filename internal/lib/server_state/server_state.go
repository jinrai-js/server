package server_state

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jinrai-js/server/internal/lib/app_error"
	"github.com/jinrai-js/server/internal/lib/interfaces"
	"github.com/jinrai-js/server/internal/lib/redirect"
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
			if to, exists := value.(map[string]any)["redirect"]; exists {
				if v := to.(string); v != "" {
					redirect.Create(v)
				}
			}

			s.State[currentKey] = value
			return value, true
		}
	}

	return nil, false
}

func (s *State) ExportScript(ctx context.Context) string {
	export, _ := json.Marshal(map[string]any{
		"state": s.JoinStates(),
	})

	err := app_error.Get(ctx)
	var error_message string

	if err.Exists {
		error_message = fmt.Sprintf(`<!-- ############(%s)######### -->`, err.Message)
	}

	result := fmt.Sprintf(`%s<script>window.__appc__ = %s</script>`, error_message, string(export))

	return result
}

func (s *State) JoinStates() map[string]any {
	result := s.State

	if s.AppStates != nil {
		for key, value := range s.AppStates.GetWithoutSource() {
			result[key] = value
		}
	}

	return result
}
