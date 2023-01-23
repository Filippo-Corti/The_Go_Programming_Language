package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string
	var prev rune
	fmt.Scan(&input)
	for _, curr := range input {
		if isEnglishAsciiLetter(prev) && isAsciiDigit(curr) {
			fmt.Println(string(prev))
		}
		prev = curr
	}
}

func isAsciiDigit(r rune) bool {
	return isAscii(r) && unicode.IsDigit(r)
}

func isEnglishAsciiLetter(r rune) bool {
	return isAscii(r) && unicode.IsLetter(r)
}

func isAscii(r rune) bool {
	return r >= rune(0) && r <= rune(127)
}