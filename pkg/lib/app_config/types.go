package app_config

import "github.com/jinrai-js/go/pkg/lib/state_interface"

type Server struct {
	Dist      string
	ConfigDir string

	Components map[string]func(props any) string
	Meta       *string
	Rewrite    *func(string) string
	Proxy      *map[string]string
	Chashing   *[]string

	Assets  *bool
	Verbose bool
}

type Route struct {
	Id      int                                        `json:"id"`
	Mask    string                                     `json:"mask"`
	Content Content                                    `json:"content"`
	State   *map[string]state_interface.StateInterface `json:"state"`
}

type JsonConfig struct {
	Routes []Route           `json:"routes"`
	Proxy  map[string]string `json:"proxy"`
	Meta   string            `json:"meta"`
}

type Content []struct {
	Type         string         `json:"type"`
	TemplateName string         `json:"content,omitempty"` // html
	Key          string         `json:"key,omitempty"`     // value
	Data         Content        `json:"data,omitempty"`    // array
	Name         string         `json:"name,omitempty"`    // custom
	Props        map[string]any `json:"props,omitempty"`   // custom
}
