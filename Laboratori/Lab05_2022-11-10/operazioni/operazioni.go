package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	s, p, d := operazioni(n1, n2)
	fmt.Printf("Somma %d. Prodotto %d. Differenza %d\n", s, p, d)
}

/*
Scrivi una funzione operazioni(n1, n2 int) (int, int, int)
che accetta due interi e restituisce somma, prodotto e differenza.
*/
func operazioni(n1, n2 int) (int, int, int) {
	return n1 + n2, n1 * n2, n1 - n2
}

// func operazioni(n1, n2 int) (s int, p int, d int) {
// 	s = n1 + n2
// 	p = n1 * n2
// 	d = n1 - n2
// 	return
// }