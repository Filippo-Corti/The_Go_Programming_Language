// Scrivere un programma es09.go che legge un numero n e stampa i numeri tra 1 e n che sono dei quadrati.
package main

import "fmt"
import "math"

func main() {
	var n int
	fmt.Print("Un numero: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		_, dec :=  math.Modf(math.Sqrt(float64(i)))
		if dec == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}