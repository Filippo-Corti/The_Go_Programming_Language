/*
Scrivete un programma conta_cifre.go dotato di:

- una funzione `main` che legge una sequenza di stringhe da standard input fino a incontrare 
la parola "stop", analizza una stringa alla volta utilizzando la funzione `contaCifre` 
(vedi sotto) e alla fine stampa:
	- quante stringhe contengono almeno una cifra
	- per ogni cifra (0, 1, ..., 9), il numero di volte che quella cifra appare nella sequenza 
	di stringhe
- una funzione
     contaCifre(s string, numCifre *[10]int) (haCifre bool)
che, data una stringa, aggiorna il conteggio delle cifre incontrate e restituisce true 
se la stringa contiene almeno una cifra, false altrimenti.

Nota. Le stringhe possono contenere caratteri di qualsiasi tipo (cifre, lettere, simboli, 
	punteggiatura, lettere accentate, ecc.).

Il programma NON deve fare uso di variabili globali.

Domanda: che prototipo (signature) deve avere la funzione `contaCifre`?
- ha parametri? se sì, quanti, di che tipi e con che finalità?
- restituisce valori? se sì, quanti, di che tipi e con che finalità?
Soffermatevi su questi punti per progettare la funzione prima di scriverla.

Esempio di esecuzione:

$ go run conta_cifre.go
c140 c140
c0m3 t1 ch14m1?
bye bye
stop
5 stringhe con cifre
[0 1 2 3 4 5 6 7 8 9]
[3 5 0 1 3 0 0 0 0 0]

*/

package main

import (
	"fmt"
	//"log"
	//"os/exec"

	//"strings"
	"testing"
	//"bytes"
	//"io"
	//"os"
)

var prog = "./conta_cifre"

func TestContaCifreBoolean(t *testing.T) {
	var numCifre [10]int

	st := "ciao65656565"
	if !contaCifre(st, &numCifre) {
		t.Fail()
	}
}

func TestContaCifre(t *testing.T) {
	var numCifre [10]int

	st := "ciao65656565alkdjaslj lkj lakjl dskja98989898"
	if !contaCifre(st, &numCifre) {
		t.Fail()
	}

	fmt.Println(numCifre)

	if numCifre[0] != 0 || numCifre[5] != 4 {
		t.Fail()
	}
}

func TestMainEsempio(t *testing.T) {
	LanciaGenerica(t,
		prog,
		"c140 c140\nc0m3 t1 ch14m1?\nbye bye\nstop",
		"5 stringhe con cifre\n[0 1 2 3 4 5 6 7 8 9]\n[3 5 0 1 3 0 0 0 0 0]\n",
		"NIENTE")
}

func TestMainPiccolo(t *testing.T) {
	LanciaGenerica(t,
		prog,
		"senzaCifre\nQUI inv3c3 c'è\nstop",
		"1 stringhe con cifre\n[0 1 2 3 4 5 6 7 8 9]\n[0 0 0 2 0 0 0 0 0 0]\n",
		"NIENTE")
}
