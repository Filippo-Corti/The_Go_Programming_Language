package main

import (
	"fmt"
	"os"
)

func repInPlace(stringa []rune, old rune, new rune) {
	for i := 0; i < len(stringa); i++ {
		if stringa[i] == old {
			stringa[i] = new
		}
	}
}

func main() {
	stringa := []rune(os.Args[1])
	repInPlace(stringa, rune(os.Args[2][0]), rune(os.Args[3][0]))
	fmt.Println(string(stringa))
}