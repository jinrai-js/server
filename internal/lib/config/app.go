package config

import "github.com/jinrai-js/server/internal/lib/interfaces"

type App struct {
	Content Content
	States  interfaces.States
}
