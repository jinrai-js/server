package context

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jinrai-js/go/internal/components"
	"github.com/jinrai-js/go/internal/tools"
)

func (r Context) GetHTML(content Content, keys []string) string {
	var result string

	for _, props := range content {
		switch props.Type {
		case "html":
			template := tools.GetTemplate(r.OutDir, props.TemplateId)
			template = r.customReplace(template, keys)
			template = r.templateReplace(template, keys)
			result += template

		case "array":
			list := r.mapByKeys(func(key string) string {
				return r.GetHTML(props.Content, append([]string{key}, keys...))
			}, props.ContentKey, keys)

			result += strings.Join(list, "###")
		}
	}

	return result
}

func (r Context) mapByKeys(callback func(key string) string, path string, keys []string) []string {
	var result []string

	for key := range r.getSliceByPath(path, keys) {
		result = append(result, callback(strconv.Itoa(key)))
	}

	return result
}

func (r Context) templateReplace(html string, keys []string) string {
	re, _ := regexp.Compile(`{{(.*?)}}`)

	result := re.ReplaceAllStringFunc(html, func(m string) string {
		matches := re.FindStringSubmatch(m)
		data := strings.Split(matches[1]+"$", "$")
		value := r.expression(data[1], r.getValueByPath(data[0], keys))
		return value
	})

	return result
}

func (r Context) customReplace(html string, keys []string) string {
	re, _ := regexp.Compile(`<custom>(.*?)<\/custom>`)

	result := re.ReplaceAllStringFunc(html, func(m string) string {
		matches := re.FindStringSubmatch(m)
		data := tools.AnyToArray(tools.StrToJson(matches[1]))

		if len(data) < 1 {
			return ""
		}

		tagName, nameOk := tools.Conv[string](data[0])
		propsData, propsOk := tools.Conv[map[string]string](data[1])
		if !nameOk || !propsOk {
			return ""
		}

		props := make(map[string]any)
		for key, value := range propsData {
			props[key] = r.getContentValue(value, keys)
		}

		result := components.Get(tagName, props)
		return result
	})

	return result
}

func (r Context) expression(expression string, value any) string {
	switch expression {
	case "inc":
		return tools.IntToStr(tools.AnyToInt(value) + 1)
	case "dec":
		return tools.IntToStr(tools.AnyToInt(value) - 1)
	case "floor":
		intVal := tools.AnyToInt(value)
		if intVal-1 < 0 {
			return "0"
		}
		return tools.IntToStr(intVal - 1)
	default:
		return tools.AnyToStr(value)
	}
}

func (r Context) getContentValue(props string, keys []string) string {
	return "[getContentValue]"
}

func (r Context) getValueByPath(path string, keys []string) any {
	split := strings.SplitN(path, "@", 2)
	sourceIndex := split[0]
	pathItems := strings.Split(split[1], "/")

	link := r.Output.Data[sourceIndex]

	for index, pathItem := range pathItems {
		if index == 0 {
			pathItem = "data"
		}

		if strings.HasPrefix(pathItem, "[ITEM=") {
			if len(keys) == 0 {
				pathItem = pathItem[6 : len(pathItem)-1]
			} else {
				pathItem = keys[0]
				keys = keys[1:]
			}

		}

		switch v := link.(type) {
		case map[string]interface{}:
			if val, ok := v[pathItem]; ok {
				link = val
			} else {
				return ""
			}
		case map[int]interface{}:
			if val, ok := v[tools.StrToInt(pathItem)]; ok {
				link = val
			} else {
				return ""
			}
		case []interface{}:
			index := tools.StrToInt(pathItem)
			if index >= 0 && index < len(v) {
				link = v[index]
			}
		default:
			return ""
		}
	}

	return link
}

func (r Context) getSliceByPath(path string, keys []string) []any {
	value := r.getValueByPath(path, keys)
	if val, ok := value.([]any); ok {
		return val
	}

	return []any{}
}
