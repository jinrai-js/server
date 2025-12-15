package jinrai

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/jinrai-js/go/pkg/lib/config"
	"github.com/jinrai-js/go/pkg/lib/server_config"
)

type Jinrai struct {
	Server config.Server
	Json   config.JsonConfig
}

const (
	Cached = ".cached"
)

func New(dist string) (Jinrai, error) {
	configDir := filepath.Join(dist, Cached)

	json, err := server_config.New(configDir)
	if err != nil {
		return Jinrai{}, err
	}

	config := Jinrai{
		Json: json,
		Server: config.Server{
			Dist:      dist,
			ConfigDir: configDir,
		},
	}

	return config, nil
}

func NewX(dist string) Jinrai {
	static, err := New(dist)
	if err != nil {
		log.Fatal(err)
	}

	return static
}

func (c *Jinrai) ServeX(port int) {
	if err := c.Serve(port); err != nil {
		log.Fatal(err)
	}
}

func (c *Jinrai) AddComponent(component string, handler func(props any) string) {
	if c.Server.Components == nil {
		mapa := make(map[string]func(props any) string)
		c.Server.Components = &mapa
	}

	(*c.Server.Components)[component] = handler
}

func (c *Jinrai) SetRewrite(rewrite func(path string) string) {
	c.Server.Rewrite = &rewrite
}

func (c *Jinrai) Debug() {
	c.Server.Verbose = true
}

func (c *Jinrai) ServeAssets(assets bool) {
	c.Server.Assets = &assets
}

func (c *Jinrai) SetProxy(proxy map[string]string) {
	c.Server.Proxy = &proxy
}

func (c *Jinrai) SetStringProxy(str string) {
	c.Log("+ proxy:", str)
	var proxy = make(map[string]string)

	for _, service := range strings.Split(str, ",") {
		values := strings.Split(service, "=")
		proxy[values[0]] = values[1]
	}

	c.Server.Proxy = &proxy
}

func (c *Jinrai) SetMeta(meta string) {
	c.Server.Meta = &meta
}

func (c *Jinrai) SetChashing(chashing []string) {
	c.Server.Chashing = &chashing
}
