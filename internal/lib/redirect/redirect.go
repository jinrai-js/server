package redirect

type Redirect struct {
	To string
}

func Create(to string) {
	panic(&Redirect{to})
}
