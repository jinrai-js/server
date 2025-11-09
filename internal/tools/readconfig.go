package tools

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
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

func GetTemplate(outDir string, templateName string) string {
	path := filepath.Join(outDir, templateName+".html")

	return ReadHTML(path)
}
