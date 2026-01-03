package config

type Route struct {
	Id      int            `json:"id"`
	Mask    string         `json:"mask"`
	Content []Content      `json:"content"`
	States  map[string]any `json:"state"`
}

type LangSource struct {
	From string `json:"from"`
	Key  string `json:"key"`
}

type Lang struct {
	DefaultLang string     `json:"defaultLang"`
	LangBaseURL string     `json:"langBaseUrl"`
	Source      LangSource `json:"source"`
}

type JsonConfig struct {
	Routes []Route           `json:"routes"`
	Proxy  map[string]string `json:"proxy"`
	Meta   string            `json:"meta"`
	Lang   Lang              `json:"lang"`
}
