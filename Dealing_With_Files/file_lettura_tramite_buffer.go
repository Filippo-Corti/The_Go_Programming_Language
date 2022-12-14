package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Programma che riceve da riga di comando il nome di un file e un numero, e stampa
quel numero di byte letti dal file.
*/
func main() {
	nomeFile := os.Args[1]
	nByte, _ := strconv.Atoi(os.Args[2])

	f, err := os.Open(nomeFile)
	defer f.Close()
	if err != nil {
		fmt.Println("Errore in lettura!", err)
		os.Exit(1)
	}

	var b []byte
	b = make([]byte, nByte)

	x, err := f.Read(b)
	if err != nil {
		fmt.Println("Errore in lettura!", err)
		os.Exit(1)
	}
	fmt.Printf("Ho letto %d byte\n", x)
	fmt.Println("Byte letti:", b)
	fmt.Println("Stringa corrispondente:", string(b))
}