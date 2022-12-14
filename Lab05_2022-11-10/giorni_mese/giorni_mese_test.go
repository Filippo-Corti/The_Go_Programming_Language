/*
In un file giorni_mese.go definire una funzione

	giorniInMese(mese int) int

che, dato come parametro il numero corrispondente a un mese, restituisce il numero di giorni di quel mese (28 per febbraio; 30 per aprile, giugno, settembre, novembre; 31 per G,M,M,L,A,O,D).

Assumiamo che il numero passato come parametro sia in [1,12], quindi non facciamo controlli.

Usiamo uno switch. Quanti casi conterrà (incluso il caso default, se lo si usa)? 3 casi, 12 casi o quanti?

Scrivere inoltre una funzione

	main()

per invocare e testare la funzione giorniInMese.

La funzione main deve leggere da standard input una stringa nel formato gg-mm-aaaa (vedi funzione Atoi del pacchetto strconv) e stampare "il mese <x> ha <tot> giorni", dove x e tot sono numeri, usando la funzione giorniInMese per determinare tot. Chiamare il programma giorni_mese.go e caricarlo su upload (dopo aver lanciato i test)

Esempio di esecuzione
---------------------
Se il programma legge

09-11-2022

deve stampare

il mese 11 ha 30 giorni





nomefile giorni_mese.go
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

func TestTrentaNovembreAprilGiugnoSettembre(t *testing.T) {
	if giorniInMese(11) != 30 ||
		giorniInMese(4) != 30 ||
		giorniInMese(6) != 30 ||
		giorniInMese(9) != 30 {
		t.Fail()
	}
}

func TestVentottoUno(t *testing.T) {
	if giorniInMese(2) != 28 {
		t.Fail()
	}
}

func TestAltriTrentuno(t *testing.T) {
	if giorniInMese(1) != 31 ||
		giorniInMese(3) != 31 ||
		giorniInMese(5) != 31 ||
		giorniInMese(7) != 31 ||
		giorniInMese(8) != 31 ||
		giorniInMese(10) != 31 ||
		giorniInMese(12) != 31 {
		t.Fail()
	}
}
