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
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	mainCycle:
	for {
		for i := 0; i < 3; i++ {
			b := scanner.Scan() 
			if b == false {
				break mainCycle
			}
			line += scanner.Text() + " "
		}
		sumOfPriorities += getPriority(findDerangedItem(strings.Split(line, " ")))
		line = ""
	}
	fmt.Println("Sum of Priorities:", sumOfPriorities)
}

func findDerangedItem(backpacks []string) rune {
	for _, r := range backpacks[0] {
		if strings.ContainsRune(backpacks[1], r) && strings.ContainsRune(backpacks[2], r) {
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
