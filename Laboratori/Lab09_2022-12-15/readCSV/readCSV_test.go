/*

Scrivere un programma `readCSV.go` che legga da stdin una sequenza di dati in formato CSV (Comma Separated Values) così organizzati:

	timestamp,temperatura,umidità

Dove:
- `timestamp` è una stringa nel formato YYYY-MM-DD/HH:mm (anno mese giorno ora minuto)
- `temperatura` è un float
- `umidità` è un float
- il separatore è una virgola ","

Il programma deve usare Scanf per leggere ogni riga dell'input.

Il programma deve calcolare il massimo e il minimo dei valori di temperatura e umidità e stamparli con il timestamp (in formato diverso dall'originale, deve essere nella forma "X(°/%) misura delle ore HH, minuti mm del giorno DD del mese MM dell'anno YYYY") in cui sono stati misurati.

Ad esempio, il file fornito come input genera questo risultato:

minTemp: 1.0° misura delle ore 4, minuti 22 del giorno 11 del mese 12 dell'anno 2022
maxTemp: 15.0° misura delle ore 9, minuti 22 del giorno 11 del mese 12 dell'anno 2022
minHumid: 49.0% misura delle ore 9, minuti 24 del giorno 11 del mese 12 dell'anno 2022
maxHumid: 91.0% misura delle ore 9, minuti 31 del giorno 6 del mese 12 dell'anno 2022

*/

package main

import (
	//"fmt"
	"testing"
)

var prog = "./readCSV"

func TestMain(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"temp.csv",
		"temp.out",
		"NIENTE")
}

func TestScanfCalled(t *testing.T) {
	if !checkIfUsed(prog, "fmt.Scanf") {
		t.Fail()
	}
}
