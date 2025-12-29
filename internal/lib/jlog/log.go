package jlog

import "fmt"

var (
	Active = false
)

func Write(a ...any) {
	if !Active {
		return
	}

	fmt.Print(a...)
}

func Writeln(a ...any) {
	if !Active {
		return
	}

	fmt.Println(a...)
}
