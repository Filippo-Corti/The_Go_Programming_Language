package main

/* --- Day 3: Rucksack Reorganization --- */

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var sumOfPriorities int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		sumOfPriorities += getPriority(findDerangedItem(line))
	}
	fmt.Println("Sum of Priorities:", sumOfPriorities)
}

func findDerangedItem(s string) rune {
	firstHalf := s[: len(s) / 2]
	secondHalf := s[len(s) / 2 :]
	for _, r := range firstHalf {
		if strings.ContainsRune(secondHalf, r) {
			return r
		}
	}
	return '-'
}

/*
Lowercase item types a through z have priorities 1 through 26.
Uppercase item types A through Z have priorities 27 through 52.
*/
func getPriority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r - 'a') + 1
	} 
	if r >= 'A' && r <= 'Z' {
		return int(r - 'A') + 27
	} 
	return -1
}
