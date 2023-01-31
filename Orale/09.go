/*
9. Funzione che data una stringa (anche non-ascii) restituisce quanti caratteri non sono lettere latine minuscole
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(contaNonLettereLatineMinuscole("nfueas15n3èèè1."))
}

func isLetteraLatinaMinuscola(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func contaNonLettereLatineMinuscole(str string) (c int) {
	for _, lettera := range str {
		if !isLetteraLatinaMinuscola(lettera) {
			c++
		}
	}
	return c
}