package server_error

import "sync"

type server_panic_container struct{}

var mu sync.Mutex

type ExportError struct {
	Message string `json:"message"`
}

var exportList []ExportError

func Create(err error) {
	mu.Lock()
	defer mu.Unlock()
	exportList = append(exportList, ExportError{
		Message: err.Error(),
	})

	panic(server_panic_container{})
}

func Export() []ExportError {
	return exportList
}

func Catch() {
	if err := recover(); err != nil {
		if _, ok := err.(server_panic_container); ok {
			return
		}

		panic(err)
	}
}
