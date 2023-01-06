package main

import (
	"fmt"
	"os"
)

func occorrenze(s string, r rune) int {
	if len(s) == 0 {
		return 0
	}
	if rune(s[0]) == r {
		return 1 + occorrenze(s[1:], r)
	} 
	return occorrenze(s[1:], r)
}

func main() {
	fmt.Println(occorrenze(os.Args[2], rune(os.Args[1][0])))
}