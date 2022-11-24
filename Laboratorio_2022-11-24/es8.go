package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

//Leggo parole e stampo le parole con lunghezza maggiore alla media
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	testo := ""
	lenTotale := 0
	//Leggi parole
	for scanner.Scan() {
		line := scanner.Text()
		testo += line
		lenTotale += len(strings.ReplaceAll(line, " ", ""))
	}
	parole := strings.Split(testo, " ")
	media := lenTotale / len(parole)
	fmt.Println("Media:", media)
	//Stampa parole
	for i := 0; i < len(parole); i++ {
		if media < len(parole[i]){
			fmt.Println(parole[i])
		}
	}
}