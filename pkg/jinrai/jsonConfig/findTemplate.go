package jsonConfig

import (
	"log"
	"net/url"
	"regexp"
)

func (c Config) FindTemplateAndRender(url *url.URL) *Route {

	for _, route := range c.Json.Routes {
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
