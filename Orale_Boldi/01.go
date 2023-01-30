/*
1. Funzione che data una slice di Studente (nome - cognome - voto), restituisca una slice con gli studenti promossi (voto >= 18)
*/

package main

import (
	"fmt"
)

type Studente struct {
	nome, cognome string
	voto int
}

func main() {
	studenti := []Studente{Studente{"Filippo", "Corti", 15}, Studente{"Paolo", "Boldi", 30}}
	fmt.Println(studentiPromossi(studenti))
}

func studentiPromossi(studenti []Studente) (promossi []Studente) {
	for _, studente := range studenti {
		if studente.voto >= 18 {
			promossi = append(promossi, studente)
		}
	}
	return promossi
}