package config

type Server struct {
	Dist      string
	ConfigDir string

	Meta     *string
	Rewrite  *func(string) string
	Proxy    *map[string]string
	Chashing *[]string

	Assets  *bool
	Verbose bool
}

type Content []struct {
	Type         string         `json:"type"`
	TemplateName string         `json:"content,omitempty"` // html
	Key          string         `json:"key,omitempty"`     // value
	Data         Content        `json:"data,omitempty"`    // array
	Name         string         `json:"name,omitempty"`    // custom
	Props        map[string]any `json:"props,omitempty"`   // custom
}

type JsonConfig struct {
	Routes []Route           `json:"routes"`
	Proxy  map[string]string `json:"proxy"`
	Meta   string            `json:"meta"`
}
