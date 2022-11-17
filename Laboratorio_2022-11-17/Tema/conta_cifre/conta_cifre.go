package main

import  "fmt"
import "bufio"
import "os"
import "unicode"
import "strings"

/* una funzione `main` che legge una sequenza di stringhe da standard input fino a incontrare 
la parola "stop" */
func main() {
	numStringheConCifre := 0
	var contatoreCifre [10]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "stop" {
			break
		}
		//In questo modo controllo le stringhe con spazi parola per parola e non come una stringa unica
		paroleInserite := strings.Split(line, " ")
		for i := 0; i < len(paroleInserite); i++ {
			if contaCifre(paroleInserite[i], &contatoreCifre) {
				numStringheConCifre++
			}
		}
	}
	fmt.Printf("%d stringhe con cifre\n", numStringheConCifre)
	fmt.Println("[0 1 2 3 4 5 6 7 8 9]")
	fmt.Println(contatoreCifre)
}

func contaCifre(s string, numCifre *[10]int) (haCifre bool) {
	for _, r := range s {
		if unicode.IsNumber(r) {
			haCifre = true
			numCifre[r - '0']++
		}
	}
	return
}