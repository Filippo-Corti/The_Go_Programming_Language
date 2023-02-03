/* Genera un testo a caso basandosi su un testo di riferimento letto da un file il cui nome
è passato da riga di comando al programma */
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const STARTING_WORD = "mai" //Starting point per il nuovo testo generato
const TEXT_LENGTH = 20

func main() {
	rand.Seed(time.Now().UnixNano())
	testo := readFile(os.Args[1])
	dizionario := generaDizionario(testo)
	fmt.Println(generaTesto(dizionario, TEXT_LENGTH))
}

func generaTesto(dizionario map[string][]string, numeroParole int) (testo string) {
	word := STARTING_WORD
	testo += word
	for i := 0; i < numeroParole; i++ {
		newWord := dizionario[word][rand.Intn(len(dizionario[word]))]
		testo += " " + newWord
		word = newWord
	}
	return
}

func generaDizionario(testo string) map[string][]string {
	dizionario := make(map[string][]string) 
	parole := strings.Fields(testo)
	for i := 0; i < len(parole) - 1; i++ {
		current := parole[i]
		next := parole[i + 1]
		dizionario[current] = append(dizionario[current], next)
	}
	return dizionario
}

func readFile(nomeFile string) (testo string) {
	file, _ := os.Open(nomeFile)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		testo += scanner.Text()
	}
	return
}