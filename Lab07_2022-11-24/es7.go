package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"sort"
)

//Leggo parole e stampo in ordine alfabetico 
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	testo := ""
	//Leggi parole
	for scanner.Scan() {
		line := scanner.Text()
		testo += line
	}
	parole := strings.Split(testo, " ")
	//Ordina parole
	sort.Strings(parole)
	//Stampa parole
	for i := 0; i < len(parole); i++ {
		fmt.Println(parole[i])
	}
}