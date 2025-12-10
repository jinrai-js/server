package app_config

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
	Id      int                        `json:"id"`
	Mask    string                     `json:"mask"`
	Content Content                    `json:"content"`
	State   *map[string]StateInterface `json:"state"`
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

type StateInterface struct {
	Options StateOption `json:"options"`
	Key     any         `json:"key"`
}
