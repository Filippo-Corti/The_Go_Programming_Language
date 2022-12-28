package main

import (
	"fmt"
	"os"
	"sort"
)

func startsAndEndsWithTheSameLetter(s string) bool {
	return s[0] == s[len(s) - 1]
}

//Returns a map with the occurencies of the substrings with len >= 3 and the same char as first and last
func getSubstrings(s string) (m map[string]int) {
	m = make(map[string]int)
	for i := 0; i < len(s); i++ {
		for j := i + 2; j < len(s); j++ {
			substring := s[i : j + 1]
			if startsAndEndsWithTheSameLetter(substring) {
				m[substring]++
			}
		}
	}
	return
}

func printMap(m map[string]int) {
	keys := getSortedKeys(m)
	for _, key := range keys {
		fmt.Printf("%s -> Occorrenze: %d\n", key, m[key])
	}
}

func getSortedKeys(m map[string]int) (slice []string) {
	for k, _ := range m {
		slice = append(slice, k)
	}
	sort.Slice(slice, func(i, j int)bool {
		return len(slice[i]) > len(slice[j])
	})
	return
}

func main() {
	substrings := getSubstrings(os.Args[1])
	printMap(substrings)
}