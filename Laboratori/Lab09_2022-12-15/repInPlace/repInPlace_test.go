/*

Scrivere un programma `repInPlace.go` che definisca una funzione:
	func repInPlace(stringa []rune, old rune, new rune)
che modifichi la `stringa` passata sostituendo ogni occorrenza della runa `old` con la runa `new`

Aggiungere inoltre un main che accetti tre parametri a linea di comando, rispettivamente:
- stringa da modificare
- carattere sostituendo
- carattere di rimpiazzo

*/

package main

import (
	"fmt"
	"testing"
)

func TestMinimo(t *testing.T) {
	s := []rune("ciao")
	fmt.Println(string(s))
	repInPlace(s, 'c', 'C')
	fmt.Println(string(s))
	if s[0] != 'C' {
		t.Fail()
	}
}

func TestDoppio(t *testing.T) {
	s := []rune("ciaopippo")
	fmt.Println(string(s))
	repInPlace(s, 'c', 'C')
	repInPlace(s, 'p', 'Q')
	fmt.Println(string(s))
	if s[0] != 'C' {
		t.Fail()
	}
	if s[4] != 'Q' || s[6] != 'Q' || s[7] != 'Q' {
		t.Fail()
	}
}
