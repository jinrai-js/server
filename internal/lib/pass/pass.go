package pass

type container struct {
}

func Exit() {
	panic(container{})
}

func Catch() {
	if err := recover(); err != nil {
		if _, ok := err.(container); ok {
			return
		}

		panic(err)
	}
}
