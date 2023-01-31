package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

//Leggo parole e stampo al contrario
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	testo := ""
	//Leggi parole
	for scanner.Scan() {
		line := scanner.Text()
		testo += line
	}
	//Stampa parole
	parole := strings.Split(testo, " ")
	for i := len(parole) - 1; i >= 0; i-- {
		fmt.Println(parole[i])
	}
}