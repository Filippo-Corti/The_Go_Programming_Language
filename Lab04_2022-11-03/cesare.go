package main

import "fmt"

func main() {
	var s string
	var k int
	fmt.Print("Inserire sequenza di lettere e chiave di cifratura: ")
	fmt.Scan(&s, &k)
	fmt.Println(cifraCesare(s, k))
}

/*
	Riceve in input la lettera l e la chiave di cifratura
k. Restituisce l'equivalente di l in cifratura
di Cesare con chiave k
*/
func cifraCesareLettera(l rune, k int) (x rune) {
	x = l + rune(k)
	for x > 'z' {
		x = 'a' + (x - 'z' - 1)
	}
	return
}

func cifraCesare(s string, k int) (x string) {
	for _, r := range s {
		x += string(cifraCesareLettera(r,k))
	}
	return
}