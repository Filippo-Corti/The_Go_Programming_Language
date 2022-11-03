package main

import "fmt"

func main() {
	var n int
	fmt.Print("Altezza: ")
	fmt.Scan(&n)
	disegnaSlash(n)
}

func disegnaSlash(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			//Stampa tot spazi
			fmt.Print(" ")
		}
		//Stampa * e va a capo
		fmt.Println("*")
	}
}