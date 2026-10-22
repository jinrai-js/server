package jinrai

import (
	"log"

	"github.com/jinrai-js/server/internal/lib/jinrai"
)

func New(dist string) (jinrai.Jinrai, error) {
	return jinrai.New(dist)
}

func NewX(dist string) jinrai.Jinrai {
	server, err := jinrai.New(dist)

	if err != nil {
		log.Fatal(err)
	}

	return server
}
