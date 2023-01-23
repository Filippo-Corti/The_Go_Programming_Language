package main

import (
	"fmt"
	"sort"
)

func main() {
	var sequence string
	fmt.Scan(&sequence)
	if !isMadeOfParenthesisOnly(sequence) {
		return
	}
	if isBalanced(sequence) {
		fmt.Println("bilanciata")
	} else {
		fmt.Println("non bilanciata")
	}
	subSequences := getBalancedSubstrings(sequence)
	printSortedMap(subSequences)
}

func printSortedMap(m map[string]int) {
	keys := getSortedKeys(m)
	for _, key := range keys {
		fmt.Printf("%s %d\n", key, m[key])
	}
}

func getSortedKeys(m map[string]int) (keys []string) {
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) < len(keys[j])
	})
	return
}

func getBalancedSubstrings(sequence string) map[string]int {
	subSequences := make(map[string]int)
	for i := 0; i < len(sequence); i++ {
		for j := i + 1; j < len(sequence); j++ {
			subSequence := sequence[i : j + 1]
			if isBalanced(subSequence) {
				subSequences[subSequence]++
			}
		}
	}
	return subSequences
}

func isBalanced(sequence string) bool {
	var openedParenthesis int
	for _, char := range sequence {
		if char == '(' {
			openedParenthesis++
		} else {
			openedParenthesis--
		}
		if openedParenthesis < 0 {
			return false
		}
	}
	return openedParenthesis == 0
}

func isMadeOfParenthesisOnly(str string) bool {
	for _, char := range str {
		if char != '(' && char != ')' {
			return false
		}
	}
	return true
}