package components

import "encoding/json"

type componentHandler func(props any) string

var components = make(map[string]componentHandler)

func Add[T any](name string, handler func(props T) string) {
	components[name] = func(v any) string {

		var props T
		jsonString, _ := json.Marshal(v)
		if err := json.Unmarshal(jsonString, &props); err == nil {
			return handler(props)
		}

		return ""
	}
}

func Get(name string, props any) string {
	if handler, ok := components[name]; ok {
		return handler(props)
	}

	return ""
}
