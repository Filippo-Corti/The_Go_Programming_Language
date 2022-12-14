package main

import (
	"fmt"
	"bufio"
	"os"
	"unicode"
)

const LEN_ALFABETO = 26

func main() {
	contaMinu := [LEN_ALFABETO]int{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	stringa := scanner.Text()
	aggiornaConteggi(stringa, &contaMinu)
	for i := 0; i < LEN_ALFABETO; i++ {
		fmt.Printf("%c %d\n", 'a' + byte(i), contaMinu[i])
	}
}

func aggiornaConteggi(s string, contaMinu *[LEN_ALFABETO]int) {
	for _, r := range s {
		if unicode.IsLower(r) {
			(*contaMinu)[r - 'a']++
		}
	}
}