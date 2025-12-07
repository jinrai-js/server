package jsonConfig

import (
	"path/filepath"

	"github.com/jinrai-js/go/internal/tools"
	"github.com/jinrai-js/go/pkg/jinrai/context"
)

type Route struct {
	Id      int                      `json:"id"`
	Mask    string                   `json:"mask"`
	Content context.Content          `json:"content"`
	State   map[string]context.State `json:"state"`
}

type JsonStruct struct {
	Routes []Route           `json:"routes"`
	Proxy  map[string]string `json:"proxy"`
	Meta   string            `json:"meta"`
}

type Config struct {
	OutDir string
	Json   JsonStruct
}

func New(outDir string) (Config, error) {
	config := Config{
		outDir,
		JsonStruct{},
	}
	err := tools.ReadConfig(filepath.Join(outDir, "config.json"), &config.Json)

	return config, err
}
