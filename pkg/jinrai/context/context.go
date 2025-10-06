package context

import (
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
	Type       string  `json:"type"`
	TemplateId int     `json:"template,omitempty"`
	Content    Content `json:"content,omitempty"`
	ContentKey string  `json:"contentKey,omitempty"`
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
	re, _ := regexp.Compile(`@signal\[\[(.*?)\]\]`)

	result := re.ReplaceAllStringFunc(input, func(m string) string {
		matches := re.FindStringSubmatch(m)
		data := strings.SplitN(matches[1], ":", 3)
		return context.getProps(data[0], data[1], data[2])
	})

	return strings.ReplaceAll(result, "\"null\"", "null")
}

func (context Context) getProps(propType string, value string, defValue string) string {
	switch propType {
	case "search":
		return context.getSearch(value, defValue)
	case "params":
		return context.getParams(value, defValue)
	case "request":
		return context.getRequestValue(value, defValue)
	default:
		log.Fatal("Не знаю такой тип ", propType)
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
