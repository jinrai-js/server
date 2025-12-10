package server_config

import (
	"path/filepath"

	"github.com/jinrai-js/go/internal/tools"
	"github.com/jinrai-js/go/pkg/lib/app_config"
)

func New(configDir string) (app_config.JsonConfig, error) {
	var config app_config.JsonConfig

	err := tools.ReadConfig(filepath.Join(configDir, "config.json"), &config)

	return config, err
}
