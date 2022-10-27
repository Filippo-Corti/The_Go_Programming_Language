/*
Scrivere un programma es11.go che stampa ripetutamente "ore e minuti:" per chiedere due valori h e min fino a ottenere due valori validi, cioè h in [0,23] e min in [0,59], poi stampa "valori validi"
*/
package main

import "fmt"

func main() {
	var h, min int
	for {
		fmt.Print("Ore e minuti: ")
		fmt.Scan(&h, &min)
		if h >= 0 && h <= 23 && min >= 0 && min <= 59 {
			break
		}
	}
	fmt.Println("validi validi")
}