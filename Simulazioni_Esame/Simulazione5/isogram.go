package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parola := scanner.Text()
	fmt.Println(isIsogramma(parola))
}

func isIsogramma(s string) bool {
	occorrenze := make(map[rune]int)
	//Crea mappa occorrenze
	for _, r := range s {
		if !unicode.IsSpace(r) {
			occorrenze[r]++
		}
	} 
	//Verifica che ogni lettera abbia lo stesso numero di frequenze
	nFreq := occorrenze[rune(s[0])]
	for _, v := range occorrenze {
		if v != nFreq {
			return false
		}
	}
	return true
}