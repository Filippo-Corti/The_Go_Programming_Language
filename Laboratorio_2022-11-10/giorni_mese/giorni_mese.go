package main

import "fmt"
import "strings"
import "strconv"

func main() {
	var data string
	fmt.Printf("Inserire data (gg-mm-aaaa): ")
	fmt.Scan(&data)
	mese := getMese(data)
	fmt.Printf("Il mese %d ha %d giorni\n", mese, giorniInMese(mese))
}

func getMese(data string) int {
	first := strings.Index(data, "-")
	last := strings.LastIndex(data, "-")
	n, _ := strconv.Atoi(data[first + 1:last])
	return n
}


/*
giorniInMese(mese int) int
che, dato come parametro il numero corrispondente a un mese, restituisce il numero 
di giorni di quel mese (28 per febbraio; 30 per aprile, giugno, settembre, novembre; 
31 per G,M,M,L,A,O,D). Assumiamo che il numero passato come parametro sia in [1,12], 
quindi non facciamo controlli.
*/
func giorniInMese(mese int) int {
	switch mese {
	case 2: 
		return 28
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}
