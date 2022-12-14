// Scrivere un programma es04b.go che genera K = 10 numeri casuali in [0,10], conta quanti sono pari e stampa questo risultato.
package main

import "fmt"
import "time"
import "math/rand"

func main() {
	rand.Seed(time.Now().Unix())
	const K = 10
	var n, c int
	for i := 0; i < K; i++ {
		n = rand.Intn(11)
		if n % 2 == 0 {
			c++
		}
	}
	fmt.Println("Numeri pari:", c, " / 10")
}