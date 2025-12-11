package server_state

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jinrai-js/go/pkg/lib/state_interface"
)

type State struct {
	State          map[string]any
	Proxy          map[string]string
	StateInterface *map[string]state_interface.StateInterface
}

func New(proxy map[string]string, stateInterface *map[string]state_interface.StateInterface) State {
	return State{
		make(map[string]any),
		proxy,
		stateInterface,
	}
}

func (s *State) Get(ctx context.Context, stateName string, keys []string) (any, bool) {
	stateInterface := s.getStateInterface(stateName)

	// currentKey ключ данных в state с учетом keys
	// (ключ ссылается на данные из конкретного запроса)
	currentKey := stateInterface.GetCurrentKey(ctx, keys)

	if value, exists := s.State[currentKey]; exists {
		return value, true
	}

	if value, exists := stateInterface.GetValue(ctx); exists {
		s.State[currentKey] = value
		return value, true
	}

	return nil, false
}

func (s *State) getStateInterface(stateName string) *state_interface.StateInterface {
	if stateInterface, exists := (*s.StateInterface)[stateName]; exists {
		return &stateInterface
	}
	return nil

	// return state_interface.GenerateKey(ctx, &stateInterface, keys)
}

func (s *State) Export() string {
	export, _ := json.Marshal(s.State)
	result := fmt.Sprintf(`<script>window.next_f = %s</script>`, string(export))

	return result
}
