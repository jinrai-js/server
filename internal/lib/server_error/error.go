package server_error

type server_panic_container struct{}

type ExportError struct {
	Message string `json:"message"`
}

func Create(err error) {

	panic(server_panic_container{})
}

func Export() []ExportError {
	// return exportList
	return []ExportError{}
}

func Catch() {
	if err := recover(); err != nil {
		if _, ok := err.(server_panic_container); ok {
			return
		}

		panic(err)
	}
}
