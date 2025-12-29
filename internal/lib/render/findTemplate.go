package render

import (
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/jinrai-js/server/internal/lib/app_state"
	"github.com/jinrai-js/server/internal/lib/config"
	"github.com/jinrai-js/server/internal/lib/interfaces"
)

func FindTemplateAndRender(url *url.URL, routes *[]config.Route) (*[]config.Content, interfaces.States) {
	path := strings.TrimRight(url.Path, "/")

	if path == "" {
		path = "/"
	}

	for _, route := range *routes {
		re, err := regexp.Compile("^" + route.Mask + "$")
		if err != nil {
			log.Panic(err)
		}

		if !re.MatchString(path) {
			continue
		}

		state := app_state.New(route.States)

		return &route.Content, &state
	}

	return nil, nil
}
