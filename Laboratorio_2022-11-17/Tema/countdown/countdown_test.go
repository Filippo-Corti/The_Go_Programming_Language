/*
Scrivere un programma timer.go dotato di:

- una struttura `Clock` con tre campi `hour`, `min`, `sec` di tipo `int`,
 dichiarati in quest'ordine

- una funzione `countdown` (a cui passare un puntatore ad un Clock) che fa 
scalare l'orario di un secondo, invocando opportunamente `updateMin` (vedi sotto) 
quando sono da modificare anche i minuti

- una funzione `updateMin` (idem) che fa scalare l'orario di un minuto, 
invocando opportunamente `updateHour` (vedi sotto) quando sono da modificare anche le ore

- una funzione `updateHour` (idem) che fa scalare l'orario di un'ora

Stabilite voi la segnatura adeguata per le funzioni qui sopra.

La funzione `main` chiede l'orario di partenza nel formato ore minuti secondi e 
fa partire il countdown, stampando l'orario a ogni secondo, fino a raggiungere 0 0 0.

Nota. Il programma deve creare un solo Clock e aggiornarne via via i valori dei campi, 
non deve creare un nuovo Clock a ogni secondo.

**Suggerimento**: usare l'istruzione `time.Sleep(time.Duration(1) * time.Second)` 
(potete copiarla così come è) per far passare (circa) un secondo prima di stampare ogni nuovo 
orario.


Esempio

$ go run countdown.go
time (hh mm ss): 1 0 3
{1 0 2}
{1 0 1}
{1 0 0}
{0 59 59}
{0 59 58}
{0 59 57}
{0 59 56}
{0 59 55}
...
{0 0 0}

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

var prog = "./countdown"

func TestEsistenzaFunzioni(t *testing.T) {
	var timer Clock
	countdown(&timer)
	updateMin(&timer)
	updateHour(&timer)

}

/** in realtà serve solo a testare esistenza struct */
func TestZero(t *testing.T) {
	var timer Clock

	if timer.hour != 0 {
		t.Fail()
	}
}

func TestZeroDown(t *testing.T) {
	var timer Clock
	updateHour(&timer)
	fmt.Println(timer)

	if timer.hour != -1 {
		t.Fail()
	}
}

/*
func TestZeroPositive(t *testing.T) {
	var timer Clock
	//timer.min = 2
	changeHour(&timer)
	fmt.Println(timer)

	if nonZeroPositive(&timer) {
		t.Fail()
	}
}
*/

func TestNormal(t *testing.T) {
	var timer Clock
	timer.hour = 3
	timer.min = 2
	timer.sec = 20
	countdown(&timer)
	fmt.Println(timer)

	if timer.hour != 3 && timer.min != 2 && timer.sec != 19 {
		t.Fail()
	}
}

func TestScattaMin(t *testing.T) {
	var timer Clock
	timer.hour = 3
	timer.min = 2
	timer.sec = 0
	countdown(&timer)
	fmt.Println(timer)

	if timer.hour != 3 && timer.min != 1 && timer.sec != 59 {
		t.Fail()
	}
}

func TestScattaOra(t *testing.T) {
	var timer Clock
	timer.hour = 3
	timer.min = 0
	timer.sec = 0
	countdown(&timer)
	fmt.Println(timer)

	if timer.hour != 2 && timer.min != 59 && timer.sec != 59 {
		t.Fail()
	}
}

func TestMainMinimale(t *testing.T) {
	LanciaGenerica(t,
		prog,
		"0 0 4",
		"time (hh mm ss): {0 0 3}\n{0 0 2}\n{0 0 1}\n{0 0 0}\n",
		"NIENTE")
}
