package app_state

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

type states map[string]AppState

func New(data map[string]any) states {
	var result states

	for key, value := range data {
		if val, ok := value.(AppState); ok {
			result[key] = val
		}
	}

	return result
}

// Get - получить State по названию
func (s *states) Get(name string) *AppState {
	if state, exist := (*s)[name]; exist {
		return &state
	}

	return nil
}
