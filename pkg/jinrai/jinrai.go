package jinrai

import (
	"log"
	"path"

	"github.com/jinrai-js/go/internal/server"
	"github.com/jinrai-js/go/pkg/jinrai/jsonConfig"
)

type Static struct {
	Dist       string
	Api        string
	Meta       *string
	components map[string]func(props any) string
	jsonConfig.Config
	Rewrite *func(string) string
	Assets  *bool
	Verbose bool
}

const (
	Cached = ".cached"
)

func New(dist, api string, meta *string) (Static, error) {
	jconfig, err := jsonConfig.New(path.Join(dist, Cached))
	if err != nil {
		return Static{}, err
	}

	config := Static{
		dist,
		api,
		meta,
		make(map[string]func(props any) string),
		jconfig,
		nil,
		nil,
		false,
	}

	return config, nil
}

func NewX(templates, api string, meta *string) Static {
	static, err := New(templates, api, meta)
	if err != nil {
		log.Fatal(err)
	}

	return static
}

func (c *Static) Serve(port int) error {
	return server.Run(port, c.Handler, &c.Dist)
}

func (c *Static) ServeX(port int) {
	err := server.Run(port, c.Handler, &c.Dist)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Static) AddComponent(component string, handler func(props any) string) {
	c.components[component] = handler
}

func (c *Static) Proxy(rewrite func(path string) string) {
	c.Rewrite = &rewrite
}

func (c *Static) Debug() {
	c.Verbose = true
}

func (c *Static) ServeAssets(assets bool) {
	c.Assets = &assets
}
