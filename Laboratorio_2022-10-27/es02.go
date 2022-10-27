/*
Scrivere un programma es02.go che legge K = 5 numeri e di ciascuno stampa il doppio.
*/
package main

import "fmt"

func main() {
	const K = 5
	var n int
	for i := 0; i < K; i++ {
		fmt.Print("Numero: ")
		fmt.Scan(&n)
		fmt.Println(n * 2)
	}
}