package tools

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/jinrai-js/go/internal/lib/config/app_context"
)

func ReadConfig(path string, config any) error {
	fullPath := GetBaseRoot(path)

	jsonFile, err := os.ReadFile(fullPath)
	if err != nil {
		return err
	}

	json.Unmarshal(jsonFile, config)

	return nil
}

var htmlCache = make(map[string]string)

func ReadHTML(path string) string {
	if val, ok := htmlCache[path]; ok {
		return val
	}

	fullPath := GetBaseRoot(path)

	fileContent, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatal("Не удалось прочитать файл", err)
	}

	htmlCache[path] = string(fileContent)

	return htmlCache[path]
}

func GetTemplate(ctx context.Context, templateName string) string {
	server := app_context.GetServer(ctx)

	path := filepath.Join(server.ConfigDir, templateName+".html")

	return ReadHTML(path)
}
