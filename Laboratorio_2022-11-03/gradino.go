package main

import "fmt"

/*
Il programma tiene in memoria due gradini, Gradino1 e Gradino2, in quanto la sovrapposizione dei gradini avviene per un massimo di due gradini.
I gradini vengono memorizzati dalla coppia <numero di inizio del gradino>, <lunghezza attuale del gradino>.
I gradini vengono inizializzati in base al primo input, in modo che Gradino2 preceda Gradino1.
Ad ogni inserimento viene valutato se debba essere incrementata la lunghezza. Se la lunghezza del Gradino2 non viene incrementata è perché il gradino 
è terminato. Per tale motivo Gradino2 prende i valori di Gradino1 e Gradino1 viene inizializzato al nuovo gradino, appena iniziato.
Quando un gradino termina viene valutato se la lunghezza raggiunta sia la nuova lunghezza massima.
*/
func main() {
	var n, prec, max int
	fmt.Scan(&n)
	startGradino1 := n //Partenza del Primo Gradino
	startGradino2 := startGradino1 - 1 //Partenza del Secondo Gradino
	lenGradino1 := 0 //Lunghezza del Primo Gradino
	lenGradino2 := 0 //Lunghezza del Secondo Gradino
	for {
		if n - startGradino1 <= 1 { //n == startGradino1 || n - 1 == startGradino1
			lenGradino1++
		}
		if n - startGradino2 <= 1 { //n == startGradino2 || n - 1 == startGradino2
			lenGradino2++
		} else {
			//Gradino2 diventa Gradino1 e Gradino2 viene "resettato"
			if lenGradino2 > max {
				max = lenGradino2
			}
			lenGradino2, startGradino2 = lenGradino1, startGradino1
			lenGradino1, startGradino1 = 1, n
		}
		fmt.Scan(&n)
		if n < prec {
			if lenGradino2 > max { //Controllo l'ultimo gradino
				max = lenGradino2
			}
			break
		}
		prec = n
	}
	fmt.Println("Il gradino più lungo è lungo:", max)
}