
// Esempio di esecuzione

// $ go run array.go

// array: [0 1 2 3 4]
// reverse: [4 3 2 1 0]
// switchFirstLast: [0 3 2 1 4]

package main

import "fmt"

// - una costante DIM per la dimensione dell'array, dichiarata a livello di package
const DIM = 5

// - una funzione main che inizializza a piacere un array di int di dimensione DIM 
// (ad esempio con numeri 0, 1, 2, ...) e testa le due seguenti funzioni che lavorano 
// **direttamente sull'array stesso**, senza quindi dichiarare e usare altri array. 
// Il programma stampa l'array iniziale, l'array dopo aver invocato reverse e l'array 
// dopo aver invocato switchFirstLast.
func main() {
	var array [DIM]int
	inizializzaArray(&array)
	fmt.Println(array)
	reverse(&array)
	fmt.Println(array)
	switchFirstLast(&array)
	fmt.Println(array)
}

func inizializzaArray(a *[DIM]int) {
	for i := range a {
		a[i] = i
	}
}

// - una funzione reverse che inverte l'ordine dei valori in un array di dimensione DIM, 
// mettendo il primo in fondo, il secondo in penultima posizione e così via, nell'array stesso
func reverse(a *[DIM]int) {
	len := len(a)
	for i := 0; i < len/2; i++ {
		a[i], a[len - 1 - i] = a[len - 1 - i], a[i]
	}
}

// - una funzione switchFirstLast che scambia il primo con l'ultimo dei valori in un array 
// di dimensione DIM nell'array stesso
func switchFirstLast(a *[DIM]int) {
	a[0], a[len(a) - 1] = a[len(a) - 1], a[0]
}