package jinrai_value

import (
	"context"
	"encoding/json"
	"log"
)

func Parse(ctx context.Context, data any) any {
	switch v := data.(type) {
	case []any:

		result := []any{}

		for _, value := range v {
			result = append(result, Parse(ctx, value))
		}

		return result

	case map[string]any:
		if jv, exists := v["$JV"]; exists {
			if result, err := MapToJV(jv); err == nil {
				return result.GetValue(ctx)
			}
			log.Panic("Не удалось конвертировать $JV")
		}

		result := make(map[string]any)
		for key, value := range v {
			result[key] = Parse(ctx, value)
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
