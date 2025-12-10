package render

import (
	"log"
	"net/url"
	"regexp"

	"github.com/jinrai-js/go/pkg/lib/app_config"
)

func FindTemplateAndRender(url *url.URL, routes *[]app_config.Route) *app_config.Route {

	for _, route := range *routes {
		re, err := regexp.Compile("^" + route.Mask + "$")
		if err != nil {
			log.Fatal(err)
		}

		if !re.MatchString(url.Path) {
			continue
		}

		return &route
	}

	return nil
}
