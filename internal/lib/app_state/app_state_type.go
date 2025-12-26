package app_state

import (
	"encoding/json"

	"github.com/jinrai-js/server/internal/lib/interfaces"
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
	Source *StateSource `json:"source,omitempty"`
}

type AppState struct {
	Options StateOption `json:"options"`
	Key     any         `json:"key"`
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

func (s *states) GetWithoutSource() map[string]any {
	result := make(map[string]any)

	for key, state := range *s {
		source := state.Options.Source
		if source == nil {
			result[key] = state
		}
	}

	return result
}
