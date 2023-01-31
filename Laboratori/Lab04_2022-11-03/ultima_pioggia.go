package main

import "fmt"
import "io"

func main() {
	var n int
	var ultimoGiorno int
	i := 1
	for {
		fmt.Print("Giorno ", i, ": ")
		_, err := fmt.Scan(&n)
		if err == io.EOF {
			break
		}
		if n != 0 {
			ultimoGiorno = i
		}
		i++
	}
	//Stampa risultato
	fmt.Println("\n L'ultimo giorno di pioggia è il", ultimoGiorno, "giorno")
}