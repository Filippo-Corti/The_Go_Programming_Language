/*
11. Funzione che data una slice di stringhe restituisce una stringa calcolata collegando le stringhe aggiungendo una virgola fra esse
*/

package main

import (
	"fmt"
)

func main() {
	slice := []string{"ciao", "come", "stai", "?"}
	fmt.Println(merge(slice))
}

func merge(slice []string) (merged string) {
	for _, el := range slice {
		merged += el + ","
	}
	return merged[:len(merged) - 1]
}