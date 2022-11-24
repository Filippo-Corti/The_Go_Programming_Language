package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	totale := 0.0
	for scanner.Scan() {
		line := scanner.Text()
		totale += calcolaSpesa(line)
	}
	fmt.Println(totale)
}

func calcolaSpesa(s string) float64 {
	var err error
	var prezzo, quantita, sconto float64
	dati := strings.Split(s, "#")
	prezzo, err = strconv.ParseFloat(dati[0], 64)
	if err != nil {
		return -1
	}
	quantita, err = strconv.ParseFloat(dati[1], 64)
	if err != nil {
		return -1
	}
	sconto, err = strconv.ParseFloat(dati[2], 64)
	if err != nil {
		return -1
	}
	return (prezzo * quantita) - (prezzo * quantita * sconto)
}