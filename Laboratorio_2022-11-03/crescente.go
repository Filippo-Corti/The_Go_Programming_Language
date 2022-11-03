package main

import "fmt"
import "unicode"

/*
	Scrivere un programma crescente.go che legge da standard input una stringa di sole lettere minuscole e la stampa inserendo un '-' ogni volta che una lettera viene prima in ordine alfabetico della lettera precedente.
*/
func main() {
	var s string
	fmt.Print("Inserire stringa: ")
	fmt.Scan(&s)
	fmt.Print(string(s[0]))
	for i := 1; i < len(s); i++ {
		if !(unicode.IsLetter(rune(s[i])) && unicode.IsLower(rune(s[i]))) {
			fmt.Println("\nORRORE")
			return
		}
		if s[i] < s[i - 1] {
			fmt.Print("-")
		}
		fmt.Print(string(s[i]))
	}
	fmt.Println()
}