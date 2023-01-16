package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

//I test non vanno perché mancano i file di input :(

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Fornire 1 nome di file")
		os.Exit(0)
	}
	lista, ok := parseInput(os.Args[1])
	if !ok {
		fmt.Println("File non accessibile")
		os.Exit(0)
	}
	fmt.Println(lista)
	listaCancellata := cancella(lista)
	fmt.Println(listaCancellata)
}

func parseInput(fileName string) (input []string, ok bool) {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return 
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, strings.Fields(line)...)
	}
	ok = true
	return
}

func cancella(lista []string) (out []string) {
	for i := 0; i < len(lista); i++ {
		intValue, err := strconv.Atoi(lista[i])
		if err == nil {
			i += intValue
			continue //Not a number
		}
		out = append(out, lista[i])
	}
	return out
}