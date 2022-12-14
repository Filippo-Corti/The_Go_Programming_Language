package main

import (
	"fmt"
	"os"
	"bufio"
)

/*
Programma che riceve da riga di comando il nome di un file, lo legge e lo stampa sullo
schermo una riga per volta
*/
func main() {
	nomeSorgente := os.Args[1]

	sorgente, err := os.Open(nomeSorgente)
	defer sorgente.Close()
	if err != nil {
		fmt.Println("Errore in apertura del file sorgente!", err)
		os.Exit(1)
	}

	//Lettura:
	scanner := bufio.NewScanner(sorgente)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	
}