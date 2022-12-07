package main

import (
	"fmt"
	"github.com/mnbi/gostudy"
	"github.com/mnbi/gostudy/cat"
)

func greeting() string {
	return "Hello, world!"
}

func main() {
	fmt.Printf("gostudy %v\n", gostudy.Version())

	msg := greeting()
	fmt.Println(msg)

	pathname := ""
	_ = cat.Cat(pathname)
}
