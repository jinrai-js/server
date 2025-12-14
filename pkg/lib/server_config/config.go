package server_config

import (
	"path/filepath"

	"github.com/jinrai-js/go/internal/tools"
	"github.com/jinrai-js/go/pkg/lib/config"
)

func New(configDir string) (config.JsonConfig, error) {
	var config config.JsonConfig

	err := tools.ReadConfig(filepath.Join(configDir, "config.json"), &config)

	return config, err
}
