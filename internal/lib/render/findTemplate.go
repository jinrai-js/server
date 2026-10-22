package render

import (
	"log"
	"regexp"

	"github.com/jinrai-js/server/internal/lib/app_state"
	"github.com/jinrai-js/server/internal/lib/config"
	"github.com/jinrai-js/server/internal/lib/interfaces"
)

func FindTemplateAndRender(url string, routes *[]config.Route) (*[]config.Content, interfaces.States) {
	for _, route := range *routes {
		re, err := regexp.Compile("^" + route.Mask + "$")
		if err != nil {
			log.Panic(err)
		}

		if !re.MatchString(url) {
			continue
		}

		state := app_state.New(route.States)

		return &route.Content, &state
	}

	return nil, nil
}
