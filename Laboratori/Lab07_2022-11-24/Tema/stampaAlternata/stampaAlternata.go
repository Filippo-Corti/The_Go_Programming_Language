package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	listaPari := []string{}
	listaDispari := []string{}
	cont := 1
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if cont % 2 == 0 {
			listaPari = append(listaPari, line)
		} else {
			listaDispari = append(listaDispari, line)
		}
		cont++
	}
	for _, s := range listaPari {
		fmt.Println(s)
	}
	for _, s := range listaDispari {
		fmt.Println(s)
	}
}