package main

import (
	"fmt"
	"strings"
)

func main() {
	var signal string
	fmt.Scanf("%s", &signal)
	fmt.Println(findMessageMarkerIndex(signal))
}

func findMessageMarkerIndex(s string) int {
	for i := 0; i < len(s) - 14; i++ {
		if isAMarker(s[i:i+14]) {
			return i + 14
		}
	}
	return -1
}

func isAMarker(s string) bool {
	for i, r := range s {
		if strings.ContainsRune(s[i + 1:], r) {
			return false
		}
	}
	return true
}