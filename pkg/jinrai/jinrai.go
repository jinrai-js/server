package jinrai

import (
	"log"

	"github.com/jinrai-js/go/internal/server"
	"github.com/jinrai-js/go/pkg/jinrai/jsonConfig"
)

type Static struct {
	Templates  string
	Api        string
	Meta       string
	Components map[string]func(props any) string
	jsonConfig.Config
	Rewrite *func(string) string
}

func New(templates, api, meta string, rewrite *func(string) string) (Static, error) {
	jconfig, err := jsonConfig.New(templates)
	if err != nil {
		return Static{}, err
	}

	config := Static{
		templates,
		api,
		meta,
		make(map[string]func(props any) string),
		jconfig,
		rewrite,
	}

	return config, nil
}

func NewX(templates, api, meta string, rewrite *func(string) string) Static {
	static, err := New(templates, api, meta, rewrite)
	if err != nil {
		log.Fatal(err)
	}

	return static
}

func (c Static) Serve(port int) error {
	return server.Run(port, c.Handler)
}

func (c Static) ServeX(port int) {
	err := server.Run(port, c.Handler)
	if err != nil {
		log.Fatal(err)
	}
}

func (c Static) AddComponent(component string, handler func(props any) string) {
	c.Components[component] = handler
}
