package main

import (
	"fmt"
	"os"
)

/*
Programma che riceve da riga di comando il nome di un file esistente e un altro nome, e 
fa una copia.
*/
func main() {
	nomeSorgente := os.Args[1]
	nomeDestinazione := os.Args[2]

	sorgente, err := os.Open(nomeSorgente)
	defer sorgente.Close()
	if err != nil {
		fmt.Println("Errore in apertura del file sorgente!", err)
		os.Exit(1)
	}

	destinazione, err := os.Create(nomeDestinazione)
	defer destinazione.Close()
	if err != nil {
		fmt.Println("Errore in apertura del file destinazione!", err)
		os.Exit(1)
	}

	var buffer []byte
	buffer = make([]byte, 1024)

	for {
		nLetti, _ := sorgente.Read(buffer)
		if nLetti == 0 {
			break
		}
		destinazione.Write(buffer[:nLetti])
	}
	
}