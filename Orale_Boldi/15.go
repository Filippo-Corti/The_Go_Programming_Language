/*
15. Funzione che data una slice di interi sommi gli interi in posizione pari, sommi gli interi in posizione dispari e restituisca la differenza
*/

package main

import (
	"fmt"
)

func main() {
	interi := []int{1, 3, 4, 7}
	fmt.Println(differenzaPariDispari(interi))
}

func differenzaPariDispari(interi []int) (diff int) {
	for i := 0; i < len(interi); i += 2 {
		diff += interi[i]
		diff -= interi[i + 1]
	}
	return
}