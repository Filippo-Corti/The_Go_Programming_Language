/*
12. Programma che disegna una X di asterischi data l'altezza (dispari) da linea di comando
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	disegnaX(n) // n dispari
}

func disegnaX(n int) {
	disegnaXSopra(n / 2)
	fmt.Println(strings.Repeat(" ", n / 2) + "*")
	disegnaXSotto(n / 2)
}

func disegnaXSopra(altezza int) {
	for i := 0; i < altezza; i++ {
		disegnaRiga(i, altezza * 2 - 2 * i - 1)
	}
}

func disegnaXSotto(altezza int) {
	for i := 0; i < altezza; i++ {
		disegnaRiga(altezza - i - 1, (i * 2) + 1)
	}
}

func disegnaRiga(spaziPrima int, spaziDentro int) {
	fmt.Println(strings.Repeat(" ", spaziPrima) + "*" + strings.Repeat(" ", spaziDentro) + "*")
}