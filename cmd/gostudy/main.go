package main

import (
	"fmt"
	"github.com/mnbi/gostudy"
	"github.com/mnbi/gostudy/cat"
	"os"
)

func greeting() string {
	return "Hello, world!"
}

func main() {

	cmd, args := parseArgs(os.Args)

	switch cmd {
	case "cat":
		fmt.Println("***")

		for i := 0; i < len(args); i++ {
			_ = cat.Cat(args[i])
			fmt.Println("---")
		}

		fmt.Println("***")
	case "greeting":
		msg := greeting()
		fmt.Println(msg)
	case "version":
		fmt.Printf("gostudy %v\n", gostudy.Version())
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
