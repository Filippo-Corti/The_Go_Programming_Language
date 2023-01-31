package main

import (
	"fmt"
	"os"
)

/*
Programma che riceve da riga di comando il nome di un file, lo legge tutto intero e 
lo stampa convertendolo in stringa
*/
func main() {
	nomeSorgente := os.Args[1]

	b, _ := os.ReadFile(nomeSorgente)
	s := string(b)
	fmt.Println(s)

}