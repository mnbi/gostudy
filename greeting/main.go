// Implementation for LAPL
// Cahpter 0: Hello, world!
package main

import (
	"fmt"
)

const msg = "Hello, world!"

func greeting() string {
	return msg
}

func main() {
	fmt.Println(greeting())
}

/*
   How to run the program:
     > go run ./greeting

   How to build the source code:
     > go build ./greeting

   How to write a string literal in source code:
     -> see the function, `greeting()`

   How to output a message to the screen:
     -> use `fmt.Println(...)`
*/
