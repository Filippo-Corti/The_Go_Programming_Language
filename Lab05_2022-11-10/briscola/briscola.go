package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Printf("Punti mano: %d\n", punti(s))
}

/*
punti(s string) int 
che accetta una stringa di 3 caratteri che rappresenta una mano di tre carte 
e restituisce il punteggio complessivo relativo per il gioco della briscola. 
Ad esempio per la mano "53J" restituisce 12 (10 della carta 3 + 2 del fante). 
Se la stringa contiene un simbolo che non corrisponde a nessuna carta, la funzione restituisce -1.
*/
func punti(s string) (sum int) {
	for _, r := range s {
		punti := puntiCarta(r)
		if punti == -1 {
			return -1
		}
		sum += punti
	}
	return sum
}

func puntiCarta(r rune) int {
	switch r {
	case 'A': 
		return 11
	case '3':
		return 10
	case 'K': 
		return 4
	case 'Q':
		return 3
	case 'J': 
		return 2
	case '7', '6', '5', '4', '2':
		return 0
	default:
		return -1
	}
}