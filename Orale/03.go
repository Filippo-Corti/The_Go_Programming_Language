/*
3. Funzione che data una slice di bool e un numero k calcola e restituisce l'and dei primi k booleani della slice
*/

package main

import (
	"fmt"
)

func main() {
	booleans := []bool{true, true, false, true, false}
	fmt.Println(multipleAnd(booleans, 2))
	fmt.Println(multipleAnd(booleans, 3))
}

// L'and di K booleani è true se e solo se tutti i booleani sono true
func multipleAnd(bools []bool, k int) bool {
	for i := 0; i < k; i++ {
		if !bools[i] {
			return false
		}
	}
	return true
}