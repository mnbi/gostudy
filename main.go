package main

import "fmt"

func greeting() string {
	return "Hello, world!"
}

func main() {
	msg := greeting()
	fmt.Println(msg)
}
