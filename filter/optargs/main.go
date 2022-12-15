package main

import (
	"fmt"
	"os"
)

func greeting() string {
	return "Hello, world!"
}

func main() {

	cmd, _ := parseArgs(os.Args)

	switch cmd {
	case "greeting":
		msg := greeting()
		fmt.Println(msg)
	case "help":
		appName := os.Args[0]
		fmt.Printf("%s [command] [args...]\n", appName)
		fmt.Println("command:")
		fmt.Println("  cat")
		fmt.Println("  greeting")
		fmt.Println("  version")
		fmt.Println("  help")
	default:
		fmt.Println("unknown command")
	}

}

func parseArgs(args []string) (string, []string) {
	if len(args) <= 1 {
		return "", args
	}

	return args[1], args[2:]
}
