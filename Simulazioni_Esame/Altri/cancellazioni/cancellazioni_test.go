/*
Cancellazioni
=============

Scrivere un programma 'cancellazioni.go' dotato di:

- una funzione

 	func cancella(lista []string) []string

  che, per ogni numero n >0 (intero) presente in lista,
  cancella il numero stesso e gli n elementi successivi della lista
  (o quelli che ci sono per arrivare alla fine della lista)
  e restituisce la nuova lista così prodotta;

- una funzione main() che legge da file (il cui nome viene passato
  come parametro su linea di comando) un testo di parole
  e numeri non negativi, stampa la lista stessa e poi
  la lista ottenuta cancellando, per ogni numero n presente in lista,
  il numero stesso e gli n elementi successivi (vedi sopra).

Se il programma viene lanciato con un numero di argomenti
diverso da 1, deve terminare stampando "Fornire 1 nome di file".
Se invece il file non esiste, deve terminare stampando "File non accessibile".

Esempi
------

$ go run cancellazioni.go uno.input
[uno due 2 tre quattro cinque 1 sei sette 5 otto nove dieci]
[uno due cinque sette]

$ go run cancellazioni.go due.input
[uno due 2 tre quattro cinque 1 sei sette 1 otto nove dieci]
[uno due cinque sette nove dieci]

$ go run cancellazioni.go
Fornire 1 nome di file

$ go run cancellazioni.go  FileNonEsistente.txt
File non accessibile

*/

package main

import (
	"fmt"
	"os/exec"
	"testing"
	//"strings"
	//"log"
)

func TestMainMultiplo(t *testing.T) {
	lancia(t, "[uno due 2 tre quattro cinque 1 sei sette 5 otto nove dieci]\n[uno due cinque sette]\n", "uno.input")
	lancia(t, "[uno due 2 tre quattro cinque 1 sei sette 1 otto nove dieci]\n[uno due cinque sette nove dieci]\n", "due.input")
	lancia(t, "[A b C d E f 2 A b C d E f A 3 b C d E f uno due 2 tre quattro cinque 1 sei sette 1 otto nove dieci]\n[A b C d E f C d E f A E f uno due cinque sette nove dieci]\n", "tre.input")

	lancia(t, "File non accessibile\n", "nulla")

	lancia(t, "Fornire 1 nome di file\n")
}

func lancia(t *testing.T, expected string, in ...string) {
	fmt.Println("### Questo test verifica output atteso! (presuppone che il sorgente da testare sia già stato compilato)")

	subproc := exec.Command("./cancellazioni", in...)

	stdout, stderr := subproc.CombinedOutput()

	if stderr != nil {
		t.Errorf("Fallito: %s\n", stderr)
	}

	fmt.Printf("Actual output:\n%s\n", string(stdout))

	fmt.Printf("Expected output:\n%s\n", expected)

	if string(stdout) != expected {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}

	subproc.Process.Kill()
}

func TestFunzione(t *testing.T) {
	fmt.Println("### Questo test verifica output atteso! (presuppone che il sorgente da testare sia già stato compilato)")

	var lista = []string{"uno", "due", "2", "tre", "quattro", "cinque"}
	fmt.Println(lista)

	//n:=fmt.Sprintf("%#v", cancella(lista))
	n := fmt.Sprint(cancella(lista))

	fmt.Println(n)

	if n != "[uno due cinque]" {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
}
