// Implementation for LAPL
// Chapter 1: Making a filter (count)

// Counts lines, words, runes, and bytes in the input text.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Accepts a word, then count runes in it.
func countRunes(word string) int {
	return len([]rune(word))
}

// Accepts a line of text, then count words and runes in it.  A word
// is maximal string of characters delimited by whitespaces.  It
// ignores whitespaces in counting runes.
func count(text string) (int, int) {
	var dw, dr = 0, 0

	for _, w := range strings.Fields(text) {
		dw++
		dr += countRunes(w)
	}

	return dw, dr
}

func main() {
	var (
		col int // count of lines
		cow int // count of words
		cor int // count of runes
		cob int // count of bytes
	)

	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var dw, dr = 0, 0
	var text string

	for scanner.Scan() {
		col++
		// assume a newline code is 1 byte, since almost unix uses
		// 0x0a as a newline.
		cob++

		text = scanner.Text()
		dw, dr = count(text)
		cow += dw
		cor += dr
		cob += len(text)
	}

	// use the format as same as the `wc` command (in macOS 13.1).
	fmt.Fprintf(writer, " %7d %7d %7d %7d\n", col, cow, cor, cob)
	writer.Flush()
}
