package config

type Route struct {
	Id      int            `json:"id"`
	Mask    string         `json:"mask"`
	Content Content        `json:"content"`
	States  map[string]any `json:"state"`
}
