package conventor

import (
	"encoding/json"
	"log"
)

type JV struct {
	Key       string `json:"key"`
	Type      string `json:"type"`
	Def       any    `json:"def"`
	Separator string `json:"separator"`
}

func Parse(data any) any {
	switch v := data.(type) {
	case []any:

		result := []any{}

		for _, value := range v {
			result = append(result, Parse(value))
		}

		return result

	case map[string]any:
		if jv, exists := v["$JV"]; exists {
			if result, err := MapToJV(jv); err == nil {
				return result
			}
			log.Panic("not success to convert #JV")
		}

		result := make(map[string]any)
		for key, value := range v {
			result[key] = Parse(value)
		}

		return result

	default:
		return v
	}
}

func MapToJV(m any) (JV, error) {
	var jv JV
	b, err := json.Marshal(m)
	if err != nil {
		return jv, err
	}
	err = json.Unmarshal(b, &jv)
	return jv, err
}
