package server_config

import (
	"path/filepath"

	"github.com/jinrai-js/server/internal/lib/config"
	"github.com/jinrai-js/server/internal/tools"
)

func New(configDir string) config.JsonConfig {
	var config config.JsonConfig

	tools.ReadConfig(filepath.Join(configDir, "config.json"), &config)

	return config
}
