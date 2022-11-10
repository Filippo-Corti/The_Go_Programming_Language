/*

In un file operaazioni.go definire le seguenti funzioni:

- operazioni(n1, n2 int) (int, int, int)
che accetta due interi e restituisce somma, prodotto e differenza.
Scrivine una versione con variabili di ritorno nominate e una senza (puoi commentare una delle due versioni per caricare un file solo)

Scrivi un main per invocare e testare la funzione. Il programma legge da standard input due int.

nomefile: operazioni.go

*/

package main

import (
	//"fmt"
	//"log"
	//"os/exec"

	//"strings"
	"testing"
	//"bytes"
	//"io"
	//"os"
)

func TestMinimale(t *testing.T) {
	s, p, d := operazioni(2, 3)
	if s != 5 || p != 6 || d != -1 {
		t.Fail()
	}
}

func TestNegativi(t *testing.T) {
	s, p, d := operazioni(-2, -3)
	if s != -5 || p != 6 || d != 1 {
		t.Fail()
	}
}
