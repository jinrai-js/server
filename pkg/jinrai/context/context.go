package context

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/jinrai-js/go/pkg/jinrai/serverState"
)

type Input struct {
	Path   []string
	Search url.Values
}

type Context struct {
	Input       Input
	ServerState serverState.State
	OutDir      string
}

type Export struct {
	Url    string `json:"url"`
	Input  any    `json:"input"`
	Result any    `json:"result"`
}

type Content []struct {
	Type         string         `json:"type"`
	TemplateName string         `json:"content,omitempty"` // html
	Key          string         `json:"key,omitempty"`     // value
	Data         Content        `json:"data,omitempty"`    // array
	Name         string         `json:"name,omitempty"`    // custom
	Props        map[string]any `json:"props,omitempty"`   // custom
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
		ServerState: serverState.New(make(map[string]string)),
		OutDir:      outDir,
	}
}

// func (context *Context) ExecuteRequests(proxy *map[string]string, requests Requests, rewrite *func(string) string) {
// 	for index, request := range requests {
// 		if request.Method != "POST" {
// 			log.Println("Пропускаю метод", request.Method)
// 			continue
// 		}

// 		requestPath := request.URL
// 		if rewrite != nil {
// 			requestPath = (*rewrite)(requestPath)
// 		}

// 		input := clearInput(request.Input)
// 		jsonBody := context.initProps(input)
// 		host, err := getHost(requestPath, proxy)
// 		if err != nil {
// 			log.Println("host is not defined", requestPath)
// 			continue
// 		}

// 		result, _ := tools.Post(host+requestPath, jsonBody)
// 		key := strconv.Itoa(index) + "#" + request.URL
// 		context.Output.Data[key] = result

// 		context.Output.Export = append(context.Output.Export, Export{
// 			Url:    request.URL,
// 			Input:  tools.StrToJson(jsonBody),
// 			Result: result,
// 		})
// 	}
// }

func getHost(requestPath string, proxy *map[string]string) (string, error) {
	if proxy == nil {
		return "", errors.New("proxy is not defined")
	}

	for prefix, host := range *proxy {
		if strings.HasPrefix(requestPath, prefix) {
			return host, nil
		}
	}

	return "", errors.New("not prefix in proxy")
}

func clearInput(input string) string {
	input = strings.ReplaceAll(input, "\\", "")
	input = strings.ReplaceAll(input, "\\", "")

	return input
}

func (context *Context) initProps(input string) string {
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

func (context *Context) getProps(jv JinraiValue) string {
	switch jv.Type {
	case "searchString":
		return context.getSearch(jv.Key, jv.Default)
	case "searchArray":
		log.Fatal("Не знаю как обработать searchArray")
		return ""

	case "paramsIndex": //
		return context.getParams(jv.Key, jv.Default)
	case "request":
		return context.getRequestValue(jv.Key, jv.Default)
	default:
		log.Fatal("Не знаю такой тип ", jv.Type)
		return "null"
	}
}

func (context *Context) getParams(value string, defValue string) string {
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

func (context *Context) getSearch(key string, defValue string) string {
	if context.Input.Search.Has(key) {
		return context.Input.Search.Get(key)
	}

	return defValue
}

func (context *Context) getRequestValue(value string, defValue string) string {
	return ""
}
