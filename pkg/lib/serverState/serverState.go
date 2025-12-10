package serverState

import (
	"encoding/json"
	"fmt"

	"github.com/jinrai-js/go/pkg/lib/appConfig"
)

type State struct {
	state          map[string]any
	proxy          map[string]string
	stateInterface *map[string]appConfig.StateInterface
}

func New(proxy map[string]string, stateInterface *map[string]appConfig.StateInterface) State {
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
	data, exists := (*s.stateInterface)[stateName]
	if !exists {
		return ""
	}

	return data.getStringKey(keys)
}

func (s *State) Export() string {
	export, _ := json.Marshal(s.state)
	result := fmt.Sprintf(`<script>window.next_f = %s</script>`, string(export))

	return result

}
