package main

import (
	"fmt"
	"io"
)

func main() {
	strings := readInput()
	fmt.Println(frequenze(strings))
}

func readInput() (strings []string) {
	var s string
	for {
		_, err := fmt.Scan(&s)
		if err == io.EOF {
			return 
		}
		strings = append(strings, s)
	}
	return 
}

func frequenze(s []string) map[string]int {
	frequenze := make(map[string]int)
	for _, string := range s {
		frequenze[string]++
	}
	return frequenze
}