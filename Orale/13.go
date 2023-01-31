/*
13. Funzione che data una stringa non ascii restituisce una mappa map[rune]int che ha come chiavi le rune che compaiono nella stringa e come valore le loro posizioni
    (Assumendo che non ci siano rune ripetute)
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(getPosizioniRune("123456"))
}

func getPosizioniRune(str string) map[rune]int {
	mappa := make(map[rune]int)
	for i, runa := range str {
		mappa[runa] = i
	}
	return mappa
}