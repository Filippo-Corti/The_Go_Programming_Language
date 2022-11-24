package main

import (
	"fmt"
	"unicode"
	"strings"
)

//Leggo parole e stampo senza punteggiatura
func main() {
	var testo, parola string
	for parola != "stop" {
		fmt.Scan(&parola)
		testo += parola + " "
	}
	for i := 0; i < len(testo); i++ {
		if unicode.IsPunct(rune(testo[i])) {
			testo = strings.ReplaceAll(testo, string(testo[i]), "")
		}
	}
	fmt.Println(testo)
}