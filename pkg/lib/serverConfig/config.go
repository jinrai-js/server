package serverConfig

import (
	"path/filepath"

	"github.com/jinrai-js/go/internal/tools"
	"github.com/jinrai-js/go/pkg/lib/appConfig"
)

func New(configDir string) (appConfig.JsonConfig, error) {
	var config appConfig.JsonConfig

	err := tools.ReadConfig(filepath.Join(configDir, "config.json"), &config)

	return config, err
}
