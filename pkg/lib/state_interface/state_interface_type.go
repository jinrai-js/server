package state_interface

type StateInterface struct {
	Options StateOption `json:"options"`
	Key     any         `json:"key"`
}

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
