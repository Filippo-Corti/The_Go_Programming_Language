/*
7. Funzione che data una slice di interi restituisce un'altra slice di interi che contiene gli stessi interi senza ripetizioni
*/

package main

import (
	"fmt"
)

func main() {
	slice := []int{4, 5, 6, 1, 2, 3, 4, 4, 6, 7, 8}
	fmt.Println(eliminaRipetizioni(slice))
}

func eliminaRipetizioni(vecchia []int) (nuova []int) {
	mappa := make(map[int]bool)
	for _, element := range vecchia {
		mappa[element] = true
	}
	for chiave, _ := range mappa {
		nuova = append(nuova, chiave)
	}
	return
}