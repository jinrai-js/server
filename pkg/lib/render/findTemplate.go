package render

import (
	"log"
	"net/url"
	"regexp"

	"github.com/jinrai-js/go/pkg/lib/app_state"
	"github.com/jinrai-js/go/pkg/lib/config"
)

func FindTemplateAndRender(url *url.URL, routes *[]config.Route) *config.App {

	for _, route := range *routes {
		re, err := regexp.Compile("^" + route.Mask + "$")
		if err != nil {
			log.Fatal(err)
		}

		if !re.MatchString(url.Path) {
			continue
		}

		state := app_state.New(route.States)

		return &config.App{
			route.Content,
			state,
		}
	}

	return nil
}
