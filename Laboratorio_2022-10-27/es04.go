/*
Scrivere un programma es04.go che genera K = 10 numeri casuali in [0,10] e li stampa su una riga, separati da spazi.
*/
package main

import "fmt"
import "time"
import "math/rand"

func main() {
	rand.Seed(time.Now().Unix())
	const K = 10
	for i := 0; i < K; i++ {
		fmt.Print(rand.Intn(11), " ")
	}
	fmt.Println()
}