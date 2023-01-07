package main

import (
	"fmt"
	"os"
)


func main() {
	var letta, prec string
	if len(os.Args) != 2 || !verificaValiditaParola(os.Args[1]) {
		os.Exit(1)
	}
	prec = os.Args[1]
	for {
		fmt.Println("inserisci una parola")
		fmt.Scan(&letta)
		if len(letta) != len(prec) || !verificaValiditaParola(letta) {
			fmt.Println("errore")
			continue
		}
		index, ok := areConcatenate(letta, prec)
		if ok {
			fmt.Println("corretto", index + 1, string(prec[index]), string(letta[index]))
			prec = letta
		} else {
			fmt.Println("sbagliato")
		}
	}
}

func areConcatenate(s1, s2 string) (int, bool) {
	var flag bool
	var index int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if !flag {
				flag = true
				index = i
			} else {
				return -1, false
			}
		}
	}
	return index, flag
}

func verificaValiditaParola(s string) bool {
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}