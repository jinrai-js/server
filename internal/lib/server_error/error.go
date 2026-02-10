package server_error

import (
	"net/http"

	"github.com/jinrai-js/server/internal/lib/redirect"
)

type server_panic_container struct{}

type ExportError struct {
	Message string `json:"message"`
}

func Create(err error) {

	panic(&server_panic_container{})
}

func Export() []ExportError {
	// return exportList
	return []ExportError{}
}

func Catch(w *http.ResponseWriter, r *http.Request) {
	if err := recover(); err != nil {
		switch v := err.(type) {
		case *server_panic_container:
			return

		case *redirect.Redirect:
			if w != nil {
				http.Redirect(*w, r, v.To, http.StatusMovedPermanently)
				return
			}
		}

		panic(err)
	}
}
