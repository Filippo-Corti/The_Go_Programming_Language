package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(createString(os.Args[1:]))
}

func createString(elements []string) (s string) {
	for i := 0; i < len(elements); i += 2 {
		letter := string(elements[i])
		count, _ := strconv.Atoi(elements[i + 1])
		s += strings.Repeat(letter, count)
	}
	return
}