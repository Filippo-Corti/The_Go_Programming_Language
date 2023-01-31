/*
6. Funzione che prende una mappa map[string]int e una slice []string e restituisce la somma degli interi che corrispondono alle stringhe nella slice.
    Sommare -1 per le stringhe non presenti nella mappa
*/

package main

import (
	"fmt"
)

func main() {
	chiavi := []string{"a", "b", "c"}
	mappa := make(map[string]int)
	mappa["a"] = 7
	mappa["b"] = 4
	mappa["d"] = 11
	fmt.Println(sommaPerChiavi(mappa, chiavi))
}

func sommaPerChiavi(m map[string]int, chiavi []string) (sum int) {
	for _, chiave := range chiavi {
		if v, ok := m[chiave]; ok {
			sum += v
		} else {
			sum += -1
		}
	}
	return sum
}