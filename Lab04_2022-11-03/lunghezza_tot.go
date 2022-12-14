package main

import "fmt"

func main() {
	var totLen, c int
	var s, stringhe string
	fmt.Print("Inserisci totLen: ")
	fmt.Scan(&totLen)
	for c < totLen {
		fmt.Print("Inserire stringa: ")
		fmt.Scan(&s)
		c += len(s)
		stringhe += s + " "
	}
	fmt.Println("Somma:", c)
	fmt.Println("Stringhe:", stringhe)
}