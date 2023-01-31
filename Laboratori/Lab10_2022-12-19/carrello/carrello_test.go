/*
Scrivere un programma carrello.go che simula gli spostamenti di un carrello lungo un percorso (a celle unitarie) in cui sono sono collocati degli oggetti di peso specificato. Il carrello deve avanzare e ripulire il percorso caricando gli oggetti man mano che li incontra, ma senza superare il suo carico massimo. Quando non può caricare un oggetto (per motivi di peso), deve andare all'inizio del percorso, scaricare lì (nella cella 0) il suo carico, e ripartire.

Il programma deve essere dotato di:

- una struttura Carrello con campi
	- pos int
	- carico int
  dichiarati in quest'ordine

- un medodo String() string
  per Carrello che restituisce una descrizione del carrello. Ad esempio restituisce:
  "posizione 25, carico 42"
  per un Carrello che si trova sulla cella 25 e sta trasportando 42 kg.

- una funzione aggiornaStato(c *Carrello, posizione, carico int) bool
  che aggiorna i dati del carrello (numero di cella e carico) e restuisce true, se posizione e carico non sono negativi; altrimenti non fa nessun aggiornamento e restituisce false.

- una funzione main() che legge da file, il cui nome è passato su linea di comando, la descrizione di un percorso (a celle) con oggetti da rimuovere. Una cella è delimitata da '|'; se è vuota (se c'è uno spazio bianco), in quella posizione non c'è nessun oggetto, se invece contiene un numero, in quella posizione c'è un oggetto di quel peso. Ad esempio la seguente è la descrizione di un percorso:

"| | | |12|4| | | |10| | | | |4| | | | |5| |12| | | | |3| |"

che ha un oggetto di peso 12 nella cella 3, uno di peso 4 nella cella 4, ecc. come mostrato qui sotto:

"| | | |12|4| | | |10| | | | |4| | | | |5| |12| | | | |3| |"
  0 1 2 3  4 5 6 7  8 9 10 .....

I pesi devono essere caricati su un carrello, che porta un massimo di 15 kg, e scaricati sulla prima cella del percorso.

Il main deve stampare
- il percorso come è prima di iniziare il lavoro di pulizia
- la situazione del carrello ogni volta che non può caricare un nuovo oggetto
- il percorso dopo ogni scarico
- la situazione del carrello alla fine del lavoro
- il numero di viaggi indietro fatti per rimuovere tutti gli oggetti dal percorso
- il massimo peso trovato lungo il percorso
- per ogni peso (in ordine crescente di peso), il numero di oggetti trovati di quel peso.

(Nota. si consiglia in fase di sviluppo di stampare anche la situazione del carrello dopo ogni carico)

Nel caso manchi il nome del file, il programma deve stampare il messaggio "manca nome file" e terminare.

Esempio di esecuzione per il percorso dato sopra
------------------------------------------------
| | | |12|4| | | |10| | | | |4| | | | |5| |12| | | | |3| |
carrello: posizione 3, carico 12
|12| | | |4| | | |10| | | | |4| | | | |5| |12| | | | |3| |
carrello: posizione 12, carico 14
|26| | | | | | | | | | | | |4| | | | |5| |12| | | | |3| |
carrello: posizione 19, carico 9
|35| | | | | | | | | | | | | | | | | | | |12| | | | |3| |
carrello: posizione 26, carico 15
|50| | | | | | | | | | | | | | | | | | | | | | | | | | |
carrello: posizione 0, carico 0
n viaggi: 4
peso max: 12
oggetti trovati:
1 ogg. di peso 3
2 ogg. di peso 4
1 ogg. di peso 5
1 ogg. di peso 10
2 ogg. di peso 12

*/

package main

import (
	"fmt"
	"testing"
)

var prog = "./carrello"

func TestEsisteCarrello(t *testing.T) {
	var c Carrello
	_ = fmt.Sprint(c)
}

func TestString(t *testing.T) {
	var c Carrello
	s := fmt.Sprint(c)
	if s != "posizione 0, carico 0" {
		t.Fail()
	}
}

func TestAggiorna(t *testing.T) {
	var c Carrello
	var ok bool

	aggiornaStato(&c, 5, 8)
	if c.pos != 5 || c.carico != 8 {
		fmt.Println("non aggiorna correttamente")
		t.Fail()
	}

	aggiornaStato(&c, 0, 0)
	if c.pos != 0 || c.carico != 0 {
		fmt.Println("non aggiorna correttamente")
		t.Fail()
	}

	ok = aggiornaStato(&c, -2, 3)
	if ok {
		fmt.Println("non controlla i parametri")
		t.Fail()
	}

	ok = aggiornaStato(&c, 2, -3)
	if ok {
		fmt.Println("non controlla i parametri")
		t.Fail()
	}
}

func TestMain1(t *testing.T) {
	LanciaGenericaConFileOutAtteso(
		t,
		prog,
		"NIENTE",
		"percorso.expected",
		"percorso.input")
}

func TestMain2(t *testing.T) {
	LanciaGenericaConFileOutAtteso(
		t,
		prog,
		"NIENTE",
		"percorsoBreve.expected",
		"percorsoBreve.input")
}

/*
func TestTODO(t *testing.T) {
	fmt.Println("********************** AGGIUNGERE TEST? ********************** ")
	t.Fail()
}
*/
