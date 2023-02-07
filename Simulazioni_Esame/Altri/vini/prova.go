package main

import (
	"bufio"
	"fmt"
	"os"
)

type BottigliaVino struct {
	nome     string
	anno     int
	grado    float32
	quantita int
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
	var nome string
	var anno, quantita int
	var grado float32
	_, err := fmt.Sscanf(riga, "%*[^,],%d,%f,%d", &nome, &anno, &grado, &quantita)
	fmt.Println(err)
	return BottigliaVino{nome, anno, grado, quantita}, true
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		fmt.Println(CreaBottigliaDaRiga(line))
	}
}