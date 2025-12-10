package render

import (
	"log"
	"net/url"
	"regexp"

	"github.com/jinrai-js/go/pkg/lib/appConfig"
)

func FindTemplateAndRender(url *url.URL, routes *[]appConfig.Route) *appConfig.Route {

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
