package server

import (
	"fmt"
	"net/http"
)

func Run(port int, handler func(w http.ResponseWriter, r *http.Request), assetsPath *string) error {
	http.HandleFunc("/", handler)

	if assetsPath != nil {
		fs := http.FileServer(http.Dir(*assetsPath))
		http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
