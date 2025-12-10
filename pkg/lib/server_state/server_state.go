package server_state

import (
	"encoding/json"
	"fmt"

	"github.com/jinrai-js/go/pkg/lib/app_config"
)

type State struct {
	state          map[string]any
	proxy          map[string]string
	stateInterface *map[string]app_config.StateInterface
}

func New(proxy map[string]string, stateInterface *map[string]app_config.StateInterface) State {
	return State{
		make(map[string]any),
		proxy,
		stateInterface,
	}
}

func (s *State) Get(stateName string, keys []string) (any, bool) {
	key := s.getKeyFromStateInterface(stateName, keys)

	if value, exists := s.state[key]; exists {
		return value, true
	}

	return nil, false
}

func (s *State) getKeyFromStateInterface(stateName string, keys []string) string {
	_, exists := (*s.stateInterface)[stateName]
	if !exists {
		return ""
	}

	return ""
	// return data.getStringKey(keys)
}

func (s *State) Export() string {
	export, _ := json.Marshal(s.state)
	result := fmt.Sprintf(`<script>window.next_f = %s</script>`, string(export))

	return result

}
