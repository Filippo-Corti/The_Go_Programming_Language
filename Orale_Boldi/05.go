/*
5. Funzione che dato un intero b (base) e una slice di interi (slice che rappresenta un numero in base b) restituisce il valore decimale del numero
    es: [1 0 0] con base 2 = 4
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	base := 4
	rappresentazione := []int{1, 3, 2} // = 2 + 12 + 16 = 30
	fmt.Println(converti(rappresentazione, base))
}

func converti(rappresentazione []int, base int) (convertito int, ok bool) {
	for i := len(rappresentazione) - 1; i >= 0; i-- {
		if rappresentazione[i] >= base || rappresentazione[i] < 0  {
			return
		}
		convertito += int(math.Pow(float64(base), float64(len(rappresentazione) - 1 - i))) * rappresentazione[i]
	}
	return convertito, true
}