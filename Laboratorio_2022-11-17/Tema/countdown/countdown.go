/*
Scrivere un programma timer.go dotato di:

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

import "fmt"
import _ "time"

/*
- una struttura `Clock` con tre campi `hour`, `min`, `sec` di tipo `int`,
 dichiarati in quest'ordine
*/
type Clock struct {
	hour, min, sec int
}

func main() {
	var h, m, s int
	fmt.Print("time (hh mm ss): ")
	fmt.Scan(&h, &m, &s)
	clock := Clock{h, m, s}
	for !isTimeUp(clock) {
		countdown(&clock)
		fmt.Println(clock)
		//time.Sleep(time.Duration(1) * time.Second)
	}
}

/*
- una funzione `countdown` (a cui passare un puntatore ad un Clock) che fa 
scalare l'orario di un secondo, invocando opportunamente `updateMin` (vedi sotto) 
quando sono da modificare anche i minuti
*/
func countdown(timer *Clock) {
	timer.sec--
	if timer.sec < 0 {
		timer.sec = 59
		updateMin(timer)
	}
} 

func isTimeUp(timer Clock) bool {
	return timer.hour == 0 && timer.min == 0 && timer.sec == 0
}

// - una funzione `updateMin` (idem) che fa scalare l'orario di un minuto, 
// invocando opportunamente `updateHour` (vedi sotto) quando sono da modificare anche le ore

func updateMin(timer *Clock) {
	timer.min--
	if timer.min < 0 {
		timer.min = 59
		updateHour(timer)
	}
}

// - una funzione `updateHour` (idem) che fa scalare l'orario di un'ora
func updateHour(timer *Clock) {
	timer.hour--
}