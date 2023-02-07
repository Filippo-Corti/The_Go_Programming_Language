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
  che, se i parametri sono **validi** (campi string diversi da vuoto, campi int e float maggiori di zero)
  crea una bottiglia corrispondente a questi dati e lo restituisce insieme al valore 'true',
  altrimenti restituisce una bottiglia "zero-value" e 'false'.

- una funzione main() che crea una slice di bottiglie leggendo da un file (il cui nome è passato da linea di comando)
 delle righe che contengono ognuna i dati (separati da virgole) relativi ad una bottiglia, ad es:

Rasol,2018,14,750
Camunnorum,2015,15,750
Dom Perignon,2019,12.5,1500
Balon,2013,15,750
Verdicchio,2020,11,375

  e stampa su stdout l'elenco delle bottiglie (esattamente nello stesso formato rappresentato qui sopra).
  Attenzione alle righe vuote (vedere vini.input), il programma deve essere "robusto" e ignorarle.

- una funzione
	CreaBottigliaDaRiga(riga string) (BottigliaVino, bool)
  che crea una bottiglia a partire dalla sua rappresentazione sotto forma di riga di testo, come da formato specificato sopra; se i parametri ci sono tutti e sono validi (vedi sopra),
  crea una bottiglia corrispondente a questi dati e lo restituisce insieme al valore 'true',  altrimenti restituisce una bottiglia "zero-value" e 'false'.
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
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type BottigliaVino struct {
	nome     string
	anno     int
	grado    float32
	quantita int
}

func (bottle BottigliaVino) String() string {

	output := bottle.nome + "," + strconv.Itoa(bottle.anno) + "," + strconv.FormatFloat(float64(bottle.grado), 'f', -1, 64) + "," + strconv.Itoa(bottle.quantita)

	return output

}

func CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool) {
	if nome == "" || anno <= 0 || grado <= 0 || quantita <= 0 {
		return BottigliaVino{}, false
	} else {
		return BottigliaVino{nome, anno, grado, quantita}, true
	}
}

// Rasol,2018,14,750
func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
	var nome string
	var anno, quantita int
	var gradazione float32
	fmt.Sscanf(riga, "%s,%d,%f,%d", &nome, &anno, &gradazione, &quantita)
	return CreaBottiglia(nome, anno, gradazione, quantita)
}

func AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino) {
	*cantina = append(*cantina, bot)
}

func AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino) {
	newBottle, isReal := CreaBottigliaDaRiga(bot)
	if isReal {
		AggiungiBottiglia(newBottle, cantina)
	}
}

func ContaPerTipo(nome string, cantina []BottigliaVino) (contatore int) {
	for _, bottle := range cantina {
		if bottle.nome == nome {
			contatore++
		}
	}
	return contatore
}

func main() {
	var cantina []BottigliaVino
	file, err := os.Open(os.Args[1])
	defer file.Close()
	if err != nil {
		fmt.Println("Errore nell'apertura del file:", err)
		return
	}
	scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanLines) ??????
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		AggiungiBottigliaDaRiga(line, &cantina)
		//fmt.Println(CreaBottigliaDaRiga(line))
	}

	fmt.Println(cantina)
	contatore := 0
	for _, bottle := range cantina {

		contatore += ContaPerTipo(bottle.nome, cantina)
	}

	fmt.Print("\nla cantina ha queste bottiglie", cantina, "\n il numero di bottiglie dello stesso tipo è:", contatore, "\n")

	for _, bottle := range cantina {

		fmt.Println(bottle.String())

	}

}

/*Verifica se ci sono stati errori durante la lettura
	if err := scanner.Err(); err != nil {
		fmt.Println("Errore nella lettura del file:", err)
		return
	}
}*/

/*var nome string
var anno, quantita int
var gradazione float32
if _, err := fmt.Sscanf(riga, "%s,%v,%f,%v", &nome, &anno, &gradazione, &quantita); err != io.EOF {

	fmt.Println(err)
}
return CreaBottiglia(nome, anno, gradazione, quantita)*/
