// Implementation for LAPL
// Chapter 1: Making a filter
// Ex 1.1.1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const hirakana = "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんがぎぐげござじずぜぞだぢづでどばびぶべぼぱぴぶぺぽ"
const katakana = "アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲンガギグゲゴザジズゼゾダヂヅデドバビブベボパピプペポ"

func isHirakana(kana rune) bool {
	return strings.ContainsRune(hirakana, kana)
}

func isKatakana(kana rune) bool {
	return strings.ContainsRune(katakana, kana)
}

func indexRune(s string, r rune) int {
	runes := []rune(s)
	for i, c := range runes {
		if c == r {
			return i
		}
	}
	return -1
}

func pickRune(s string, i int) rune {
	runes := []rune(s)
	return runes[i]
}

func toHirakana(kana rune) rune {
	if isKatakana(kana) {
		i := indexRune(katakana, kana)
		kana = pickRune(hirakana, i)
	}
	return kana
}

func toKatakana(kana rune) rune {
	if isHirakana(kana) {
		i := indexRune(hirakana, kana)
		kana = pickRune(katakana, i)
	}
	return kana
}

func tr(s string) string {
	original := []rune(s)
	result := make([]rune, len(original))
	_ = copy(result, original)

	for i, r := range original {
		if unicode.IsUpper(r) {
			result[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			result[i] = unicode.ToUpper(r)
		} else if isHirakana(r) {
			result[i] = toKatakana(r)
		} else if isKatakana(r) {
			result[i] = toHirakana(r)
		}
	}
	return string(result)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for scanner.Scan() {
		fmt.Fprintln(writer, tr(scanner.Text()))
		writer.Flush()
	}
}
