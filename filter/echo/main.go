// Implementation for LAPL
// Chapter 1: Making a filter
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for scanner.Scan() {
		fmt.Fprintln(writer, scanner.Text())
		writer.Flush()
	}
}

/*
   How to use the standard I/O:
     -> use `os.Stdin` and `os.Stdout`.

   How to read each line from the standard input:
     1. make a scanner, `scanner := bufio.NewWriter(os.Stdin)`,
     2. loop with `scanner.Scan()`.

   How to write a line to the standard output:
     1. make a writer, `writer := bufio.NewWriter(os.Stdout)`,
     2. use `fmt.Fprintln(writer, scanner.Text())` for each line.
*/
