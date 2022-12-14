package main

import "fmt"
import "strings"

func main() {
	var runa rune
	var str string
	
	fmt.Print("Inserire stringa: ")
	fmt.Scan(&str)
	fmt.Print("Inserire runa: ")
	fmt.Scanf("%c", &runa)
	fmt.Println("Posizione:", myIndex(str, runa))	
	fmt.Println("Posizione:", strings.Index(str, string(runa)))
}

func myIndex(s string, runa rune) int {
	for i, r := range s {
		if r == runa {
			return i
		}
	}
	return -1
}
