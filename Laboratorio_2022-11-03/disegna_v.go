package main

import "fmt"

func main() {
	var n int
	fmt.Print("Altezza: ")
	fmt.Scan(&n)
	disegnaV(n)
}

func disegnaV(n int) {
	for i := 0; i <= n; i++ {
		//Spazi
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		//Asterisco
		fmt.Print("*")
		//Spazi
		for j := 0; j < (n + n - 1) - 2 * i; j++ {
			fmt.Print(" ")
		}
		//Asterisco
		if i != n {
			fmt.Print("*") 
		}
		fmt.Println()
	}
}