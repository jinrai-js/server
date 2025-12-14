package server_state

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jinrai-js/go/pkg/lib/interfaces"
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

	// currentKey ключ данных в state с учетом keys
	// (ключ ссылается на данные из конкретного запроса)
	currentKey := appState.GetCurrentKey(ctx, keys)

	if value, exists := s.State[currentKey]; exists {
		return value, true
	}

	if value, exists := appState.GetValue(ctx); exists {
		s.State[currentKey] = value
		return value, true
	}

	return nil, false
}

func (s *State) Export() string {
	export, _ := json.Marshal(s.State)
	result := fmt.Sprintf(`<script>window.next_f = %s</script>`, string(export))

	return result
}
