package main

import "fmt"

func main() {
	var n int
	fmt.Print("Numero: ")
	fmt.Scan(&n)
	stampaFibonacci(n)
}

func stampaFibonacci(n int) {
	n1 := 0
	n2 := 1
	if n == 0 {
		return
	}
	fmt.Println("*")
	for i := 0; i < n - 1; i++ {
		f := n1 + n2
		for j := 0; j < f; j++ {
			fmt.Print("*")
		}
		fmt.Println()
		n1 = n2
		n2 = f
	}
}