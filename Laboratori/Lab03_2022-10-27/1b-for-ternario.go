package main

import "fmt"

func main() {
    /*
    Scrivere un programma che permetta di inserire 5 numeri interi e stampi, dopo il quinto numero, la somma di essi
    */
    const K = 5
	var n int
	somma := 0
	for i := 1; i <= K; i++ {
		fmt.Print("un numero: ")
		fmt.Scan(&n)
		somma = somma + n
	}
	fmt.Println(s)
}
