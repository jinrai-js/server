package serverState

import (
	"encoding/json"
	"fmt"
)

type State struct {
	state map[string]any
	proxy map[string]string
}

func New(proxy map[string]string) State {
	return State{
		make(map[string]any),
		proxy,
	}
}

func (s *State) Get(key string) (any, bool) {
	if value, exists := s.state[key]; exists {
		return value, true
	}

	return nil, false
}

func (s *State) Export() string {
	export, _ := json.Marshal(s.state)
	result := fmt.Sprintf(`<script>window.next_f = %s</script>`, string(export))

	return result

}
