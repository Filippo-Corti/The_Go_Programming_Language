package main

import "fmt"
import "strings"

func main() {
	s := "La disciplina della Human-Computer Interaction ha origini già ai tempi della Prima Rivoluzione Industriale e si occupa di progettare"
	fmt.Print("La stringa '", s, "' contiene ", contaParole(s), " parole\n")
	fmt.Print("La stringa '", s, "' contiene ", contaParoleTrim(s), " parole\n")
}

func contaParole(s string) (c int) {
	s = " " + s + " " 
	var prec rune
	for _, r := range s {
		if r == ' ' && prec != ' ' { //trovo uno spazio preceduto da un non-spazio
			c++
		}
		prec = r
	}
	return c - 1 //Sarebbe c + 1, ma aggiungo uno spazio all'inizio e alla fine quindi è come se contasse due parole in più
}

func contaParoleTrim(s string) (c int) {
	s = strings.TrimSpace(s) //Elimina extra white space a DX e SX
	var prec rune
	for _, r := range s {
		if r == ' ' && prec != ' ' { //trovo uno spazio preceduto da un non-spazio
			c++
		}
		prec = r
	}
	return c + 1
}