/*
Scrivere un programma carte.go dotato delle seguenti parti:

- definite un tipo struct Carta, per rappresentare carte da poker, con due campi, 
valore e seme (in quest'ordine), di tipo string, il primo per il valore 
("A", "1", ..., "10", "J", "Q", "K") e il secondo per il seme 
("C", "Q", "F", "P" o, volendo, i simboli dei semi "\u2665", "\u2666", "\u2663", "\u2660").

- definite 3 costanti: per il numero di semi (4), di valori (13) e di carte in un mazzo (52).

- definite una funzione 
carta(n int) (Carta, bool)
che, dato un int nell'intervallo [0,52), restituisce la corrispondente Carta da poker e true, 
dove: le prime 13 sono di cuori, in ordine dall'asso (A) al re (K);  
dalla 13 alla 25 sono di quadri; ecc.
Quindi a 10 corrisponde il fante di cuori (JC); a 13 l'asso di quadri (AQ).
Se invece l'argomento non e` nell'intervallo [0,52), restituisce false.

**Suggerimento**: sfruttare la soluzione adottata per associare numeri a nomi di
 mesi nell'esercizio sull'estrazione della data.

- definite una funzione 
estraiCarta() Carta
che genera un numero casuale in [0, 52) e restituisce la carta corrispondente. 
(Evitate di duplicare il codice della funzione carta).

- definite una funzione main che chiama estraiCarta e stampa la carta estratta.

- definite una funzione dai4Carte() che restituisce un array di 4 carte estratte con estraiCarta.

- Aggiungete al main le istruzioni per testare anche questa funzione.

**Nota**: Per generare numeri casuali, c'e` la funzione rand.Intn(n int) del pacchetto "math/rand"; per generare sequenze sempre diverse, importare il pacchetto "time" e usare l'istruzione rand.Seed(time.Now().Unix()) prima di iniziare a generare numeri casuali.
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

func TestEsistenzaFunzioni(t *testing.T) {
	estraiCarta()
	dai4Carte()
	//mazzoPoker()
	carta(4)
}
