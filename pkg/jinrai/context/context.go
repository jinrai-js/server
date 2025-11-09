package context

import (
	"encoding/json"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/jinrai-js/go/internal/tools"
)

type Input struct {
	Path   []string
	Search url.Values
}

type Output struct {
	Data   map[string]any
	Export []Export
}

type Context struct {
	Input  Input
	Output Output
	OutDir string
}

type Export struct {
	Url   string
	Body  any
	Input any
}

type Requests []struct {
	Method string `json:"method"`
	URL    string `json:"url"`
	Input  string `json:"input"`
}

type Content []struct {
	Type         string  `json:"type"`
	TemplateName string  `json:"content,omitempty"` // html
	Key          string  `json:"key,omitempty"`     // value
	Name         string  `json:"name,omitempty"`    // custom
	Props        string  `json:"props,omitempty"`   // custom
	Data         Content `json:"data,omitempty"`    // array
}

type JinraiValue struct {
	Key       string `json:"key"`
	Type      string `json:"type"`
	Separator string `json:"separator"`
	Default   string `json:"def"`
}

func New(url *url.URL, outDir string) Context {
	return Context{
		Input: Input{
			Path:   strings.Split(url.Path, "/")[1:],
			Search: url.Query(),
		},
		Output: Output{
			Data:   make(map[string]any),
			Export: []Export{},
		},
		OutDir: outDir,
	}
}

func (context Context) ExecuteRequests(host string, requests Requests, rewrite *func(string) string) {
	for index, request := range requests {
		if request.Method != "POST" {
			log.Println("Пропускаю метод", request.Method)
			continue
		}

		requestPath := request.URL
		if rewrite != nil {
			requestPath = (*rewrite)(requestPath)
		}

		jsonBody := context.initProps(request.Input)
		result, _ := tools.Post(host+requestPath, jsonBody)
		key := strconv.Itoa(index) + "#" + request.URL
		context.Output.Data[key] = result

		context.Output.Export = append(context.Output.Export, Export{
			Url:   request.URL,
			Body:  result,
			Input: tools.StrToJson(jsonBody),
		})
	}
}

func (context Context) initProps(input string) string {
	re, _ := regexp.Compile(`@JV\[\[(.*?)\]\]`)

	result := re.ReplaceAllStringFunc(input, func(m string) string {
		matches := re.FindStringSubmatch(m)

		jsonData := strings.ReplaceAll(matches[1], "\\", "")

		var jv JinraiValue
		if err := json.Unmarshal([]byte(jsonData), &jv); err != nil {
			return "null"
		}

		return context.getProps(jv)
	})

	return strings.ReplaceAll(result, "\"null\"", "null")
}

func (context Context) getProps(jv JinraiValue) string {
	switch jv.Type {
	case "searchString":
		return context.getSearch(jv.Key, jv.Default)
	case "params":
		return context.getParams(jv.Key, jv.Default)
	case "request":
		return context.getRequestValue(jv.Key, jv.Default)
	default:
		log.Fatal("Не знаю такой тип ", jv.Type)
		return "null"
	}
}

func (context Context) getParams(value string, defValue string) string {
	index, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
		return defValue
	}

	if index >= 0 && index < len(context.Input.Path) {
		return context.Input.Path[index]
	}
	return defValue
}

func (context Context) getSearch(key string, defValue string) string {
	if context.Input.Search.Has(key) {
		return context.Input.Search.Get(key)
	}

	return defValue
}

func (context Context) getRequestValue(value string, defValue string) string {
	return ""
}
