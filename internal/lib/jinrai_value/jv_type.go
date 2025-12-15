package jinrai_value

type JV struct {
	Key       string `json:"key"`
	Type      string `json:"type"`
	Def       any    `json:"def"`
	Separator string `json:"separator"`
}
