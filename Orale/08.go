/*
8. Funzione che dato un numero calcola quanti sono i suoi divisori propri (diversi da 1 e se stesso)
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(contaDivisoriPropri(700))
}

func contaDivisoriPropri(n int) (count int) {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			count++
		}
	}
	return count * 2
}