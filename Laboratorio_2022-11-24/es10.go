package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

//Leggo parole e stampo le parole con lunghezza maggiore
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	testo := ""
	maxLen := 0
	//Leggi parole
	for scanner.Scan() {
		line := scanner.Text()
		testo += line
		relativeMaxLen := getMaxLen(strings.Split(line, " "))
		if maxLen < relativeMaxLen {
			maxLen = relativeMaxLen
		}
	}
	parole := strings.Split(testo, " ")
	fmt.Println("Massimo:", maxLen)
	//Stampa parole
	for i := 0; i < len(parole); i++ {
		if len(parole[i]) == maxLen {
			fmt.Println(parole[i])
		}
	}
}

func getMaxLen(s []string) int {
	max := 0
	for _, p := range s {
		if len(p) > max {
			max = len(p)
		}
	}
	return max
}