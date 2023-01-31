/*
14. Funzione che data una stringa restituisce il numero di sottostringhe diverse che si possono creare partendo da un dato indice e fino ad un dato indice
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(contaSottostringheDiverseConIndici("cacao", 0, 3))
}

func contaSottostringheDiverseConIndici(str string, start, end int) (cont int) {
	return contaSottostringheDiverse(str[start : end + 1])
}

func contaSottostringheDiverse(str string) (cont int) {
	sottostringhe := make(map[string]bool)
	for i := 0; i < len(str); i++ {
		for j := i; j < len(str); j++ {
			sottostringhe[str[i : j + 1]] = true
		}
	}
	for _, _ = range sottostringhe { 
		cont++ 
	}
	return 
}