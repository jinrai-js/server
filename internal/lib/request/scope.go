package request

import (
	"net/url"
	"strings"
)

type Scope struct {
	Url    string
	Search url.Values
	Path   []string
}

func New(path string, searchValues url.Values) Scope {
	return Scope{
		Url:    path,
		Path:   strings.Split(path, "/")[1:],
		Search: searchValues,
	}
}
