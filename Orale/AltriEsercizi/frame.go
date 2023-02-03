package main

import (
	"fmt"
	"strings"
)

func main() {
	parole := strings.Fields("This looks better than I expected wow")
	printInFrame(parole)
}

func getMaxLength(slice []string) (max int) {
	for _, str := range slice {
		if len(str) > max {
			max = len(str)
		}
	}
	return max
}

func printInFrame(parole []string) {
	maxLength := getMaxLength(parole)
	fmt.Println(strings.Repeat("*", maxLength + 4))
	for _, parola := range parole {
		fmt.Print("* ", parola, strings.Repeat(" ", maxLength - len(parola)), " *\n")
	}
	fmt.Println(strings.Repeat("*", maxLength + 4))
}