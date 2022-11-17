/*
Scrivere un programma array.go dotato di:

- una costante DIM = 5 per la dimensione dell'array, dichiarata a livello di package

- una funzione main che inizializza a piacere un array di int di dimensione DIM (ad esempio con numeri 0, 1, 2, ...) e testa le due seguenti funzioni che lavorano **direttamente sull'array stesso**, senza quindi dichiarare e usare altri array. Il programma stampa l'array iniziale, l'array dopo aver invocato reverse e l'array dopo aver invocato switchFirstLast.

- una funzione reverse che inverte l'ordine dei valori in un array di dimensione DIM, mettendo il primo in fondo, il secondo in penultima posizione e così via, nell'array stesso

- una funzione switchFirstLast che scambia il primo con l'ultimo dei valori in un array di dimensione DIM nell'array stesso


Esempio di esecuzione

$ go run array.go

array: [0 1 2 3 4]
reverse: [4 3 2 1 0]
switchFirstLast: [0 3 2 1 4]

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

//var array = [5]int{10, 20, 30, 40, 50} // domanda da fare a loro: perché se lo metto globale i test falliscono?

func TestTestaCoda(t *testing.T) {
	var array = [5]int{10, 20, 30, 40, 50}
	switchFirstLast(&array)
	if array[0] != 50 || array[4] != 10 {
		t.Fail()
	}
}

func TestReverse(t *testing.T) {
	var array = [5]int{10, 20, 30, 40, 50}
	reverse(&array)
	if array[0] != 50 || array[1] != 40 || array[2] != 30 || array[3] != 20 || array[4] != 10 {
		t.Fail()
	}
}
