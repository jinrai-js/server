package context

type StateRequest struct {
	Input  map[string]any `json:"input"`
	Method string         `json:"method"`
	Url    string         `json:"url"`
}

type StateSource struct {
	Request StateRequest `json:"request"`
}

type StateOption struct {
	Source StateSource `json:"source"`
}

type State struct {
	Options StateOption `json:"options"`
	Key     any         `json:"key"`
}
