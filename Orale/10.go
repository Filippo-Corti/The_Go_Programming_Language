/*
10. Funzione che data una slice di stringhe restituisce true se ci sono stringhe ripetute
*/

package main

import (
	"fmt"
)

func main() {
	stringhe := []string{"a", "b", "c", "b", "d"}
	fmt.Println(hasRipetizioni(stringhe))
}


func hasRipetizioni(stringhe []string) bool {
	mappa := make(map[string]bool) 
	for _, stringa := range stringhe {
		if _, ok := mappa[stringa]; ok {
			return true
		} else {
			mappa[stringa] = true
		}
	}
	return false
}

/*
// Oppure: (peggio credo)
func contains(slice []string, str string) bool {
	for _, stringa := range slice {
		if str == stringa {
			return true
		}
	}
	return false
}

func hasRipetizioni(stringhe []string) bool {
	for i, stringa := range stringhe {
		if contains(stringhe[i+1:], stringa) {
			return true
		}
	}
	return false
}
*/