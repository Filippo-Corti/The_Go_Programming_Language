/*

Vini
----

Scrivere un programma (il cui file deve chiamarsi 'vini.go') dotato di:

- una struttura BottigliaVino con i seguenti campi (dichiarati nell'ordine):
	nome string
	anno int
	grado float32
	quantita int

- una funzione
	CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool)
  che, se i parametri sono **validi** (campi string diversi da vuoto, campi int e float maggiori di zero) crea una bottiglia corrispondente a questi dati e lo restituisce insieme al valore 'true',  altrimenti restituisce una bottiglia "zero-value" e 'false'.

- una funzione main() che crea una slice di bottiglie leggendo da un file (il cui nome è passato da linea di comando) delle righe che contengono ognuna i dati (separati da virgole) relativi ad una bottiglia, ad es:

Rasol,2018,14,750
Camunnorum,2015,15,750
Dom Perignon,2019,12.5,1500
Balon,2013,15,750
Verdicchio,2020,11,375

  e stampa su stdout l'elenco delle bottiglie (esattamente nello stesso formato rappresentato qui sopra).
  Attenzione alle righe vuote (vedere vini.input), il programma deve essere "robusto" e ignorarle.

- una funzione
	CreaBottigliaDaRiga(riga string) (BottigliaVino, bool)
  che crea una bottiglia a partire dalla sua rappresentazione sotto forma di riga di testo, come da formato specificato sopra; se i parametri ci sono tutti e sono validi (vedi sopra), crea una bottiglia corrispondente a questi dati e lo restituisce insieme al valore 'true',  altrimenti restituisce una bottiglia "zero-value" e 'false'.
  Non sono richiesti controlli sui tipi dei dati: si può assumere che i dati, se ci sono, siano nel formato corretto (ma i valori vanno controllati).

- una funzione
	AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino)
  che aggiunge una bottiglia alla lista

- una funzione
	AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino)
  che effettua la stessa operazione di AggiungiBottiglia ma partendo da una riga di testo (attenzione che è string, nel formato specificato sopra, una riga)

- una funzione
	ContaPerTipo(nome string, cantina []BottigliaVino) int
  che conta quante bottiglie dello stesso tipo (nome) di vino sono presenti nella cantina

- un **metodo** (da applicare a BottigliaVino)
    String() string
  che restituisce una riga di descrizione della bottiglia nel seguente formato:  nome,anno,grado,quantità
  (cioè ad es. restituisca "Rasol,2018,14,750" per la prima riga dell'esempio sopra

*/
package main

import (
	//"strings"
	//"log"
	//"fmt"
	//"os/exec"
	//"strings"

	"math"
	"testing"
)

func cantina() *[]BottigliaVino {
	cantina := make([]BottigliaVino, 0)
	b, _ := CreaBottiglia("Rasol", 2018, 15, 750)
	AggiungiBottiglia(b, &cantina)
	b, _ = CreaBottiglia("Camunnorum", 2015, 15, 750)
	AggiungiBottiglia(b, &cantina)
	b, _ = CreaBottiglia("Dom Perignon", 2019, 12.5, 1500)
	AggiungiBottiglia(b, &cantina)
	return &cantina
}

func TestCreaBottiglia(t *testing.T) {
	b, correct := CreaBottiglia("Dom Perignon", 2019, 12.5, 1500)
	if !correct {
		t.FailNow()
	}
	if b.nome != "Dom Perignon" {
		t.FailNow()
	}
	if !intorno(b.grado, 12.5) {
		t.FailNow()
	}
	//fmt.Println(b)
}

func TestCreaBottiglieSbagliate(t *testing.T) {
	_, correct := CreaBottiglia("", 2019, 12.5, 1500)
	if correct {
		t.FailNow()
	}
	_, correct = CreaBottiglia("Dom", 2000, -12.5, 1500)
	if correct {
		t.FailNow()
	}
	_, correct = CreaBottiglia("Dom", 2000, 12.5, -1500)
	if correct {
		t.FailNow()
	}

	_, correct = CreaBottiglia("Dom", 0, 12.5, 1500)
	if correct {
		t.FailNow()
	}

	_, correct = CreaBottigliaDaRiga("Verdicchio,2020,11")
	if correct {
		t.FailNow()
	}
}

func TestCreaBottigliaRiga(t *testing.T) {
	b, correct := CreaBottigliaDaRiga("Verdicchio,2020,11,375")
	if !correct {
		t.FailNow()
	}
	if b.nome != "Verdicchio" {
		t.FailNow()
	}
	if !intorno(b.grado, 11.0) {
		t.FailNow()
	}
	//fmt.Println(b)
}

func TestAggiungi(t *testing.T) {
	cantina := cantina() // crea
	//fmt.Println(cantina)
	//fmt.Println(len(*cantina))

	prima := len(*cantina)
	b1, _ := CreaBottiglia("Dom Perignon", 2019, 12.5, 1500)
	AggiungiBottiglia(b1, cantina)
	if len(*cantina) != (prima + 1) {
		t.FailNow()
	}
}

func TestAggiungiDaRiga(t *testing.T) {
	cantina := cantina()
	//fmt.Println(cantina)
	//fmt.Println(len(*cantina))

	prima := len(*cantina)
	AggiungiBottigliaDaRiga("Dom Perignon,2019,12.5,1500", cantina)
	if len(*cantina) != (prima + 1) {
		t.FailNow()
	}
}

func TestString(t *testing.T) {
	b, _ := CreaBottiglia("Dom Perignon", 2019, 12.5, 1500)
	//fmt.Println(b)
	if b.String() != "Dom Perignon,2019,12.5,1500" {
		t.FailNow()
	}
}

var prog = "./vini"

func TestMain(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"Rasol,2018,14,750\nCamunnorum,2015,15,750\nDom Perignon,2019,12.5,1500\nBalon,2013,15,750\nVerdicchio,2020,11,375\nRasol,2018,14,1000\nVerdicchio,2020,11,375\n",
		"vini.input")
}

func TestContaPerTipo(t *testing.T) {
	cantina := cantina() // crea
	//fmt.Println(cantina)
	//fmt.Println(len(*cantina))

	b, _ := CreaBottiglia("Rasol", 2019, 15, 750)
	AggiungiBottiglia(b, cantina)

	if ContaPerTipo("Rasol", *cantina) != 2 {
		t.FailNow()
	}
}

func intorno(a, b float32) bool {
	return math.Abs(float64(a-b)) < 10e-6
}
