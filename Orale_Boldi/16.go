/*
16. Funzione che data una slice di rune (rune ammissibili per scrivere una password) produca una password di 4 caratteri (anche ripetuti) tra quelli nella slice.
    Contare poi quante password diverse si possono creare.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(generaPasswordDi4Caratteri([]rune("abc123!&")))
	fmt.Println(contaPasswordDi4CaratteriPossibili([]rune("abc123!&")))
}

func contaPasswordDi4CaratteriPossibili(caratteriValidi []rune) (cont int) {
	passwords := make(map[string]bool)
	len := len(caratteriValidi)
	for a := 0; a < len; a++ {
		for b := 0; b < len; b++ {
			for c := 0; c < len; c++ {
				for d := 0; d < len; d++ {
					passwords[string([]rune{caratteriValidi[a], caratteriValidi[b], caratteriValidi[c], caratteriValidi[d]})] = true
				}
			}
		}
	}
	for _ = range passwords { cont++ }
	return
	// Oppure potrei eliminare le ripetizioni dai caratteriValidi e restituire 4 ^ len(caratteriValidiSenzaRipetizioni)
}

func generaPasswordDi4Caratteri(caratteriValidi []rune) (pw string) {
	for i := 0; i < 4; i++ {
		pw += string(caratteriValidi[rand.Intn(len(caratteriValidi))])
	}
	return
}

