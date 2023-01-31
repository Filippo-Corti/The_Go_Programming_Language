// Scrivere un programma es09.go che legge un numero n e stampa i numeri tra 1 e n che sono dei quadrati.
package main

import "fmt"
import "math"

func main() {
	var n int
	fmt.Print("Un numero: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		radice :=  int(math.Sqrt(float64(i)))
		if radice*radice == i {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}