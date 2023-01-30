/*
4. Funzione che data una slice di stringhe restituisce la stringa più frequente (a parità di frequenze ne restituisce una qualsiasi di priorità massima)
*/

package main

import (
	"fmt"
)

func main() {
	stringhe := []string{"ciao", "casa", "cane", "casa", "ciao", "casa", "ciao", "gatto"}
	fmt.Println(piuFrequente(stringhe))
}

func piuFrequente(stringhe []string) string {
	frequenze := make(map[string]int)
	for _, str := range stringhe {
		frequenze[str]++
	}
	return keyOfMax(frequenze)
}

func keyOfMax(m map[string]int) (max string) {
	for k, v := range m {
		if m[max] < v {
			max = k
		} 
	}
	return
}