package request

import (
	"net/url"
	"strings"
)

type Scope struct {
	Url        string
	Search     url.Values
	SearchFull string
	Path       []string
}

func New(path string, searchValues url.Values, searchFull string) Scope {
	return Scope{
		Url:        path,
		Path:       strings.Split(path, "/")[1:],
		Search:     searchValues,
		SearchFull: searchFull,
	}
}
