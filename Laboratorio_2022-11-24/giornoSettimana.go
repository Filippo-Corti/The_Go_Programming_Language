package main

import "fmt"

func main() {
	var gds string
	var n int
	giorniDellaSettimana := [7]string{"lun", "mart", "merc", "giov", "ven", "sab", "dom"}
	fmt.Println("Il primo gennaio è un: ")
	fmt.Scan(&gds)
	if arrayIndex(giorniDellaSettimana, gds) == -1 {
		fmt.Println("Input non valido")
		return
	}
	for {
		fmt.Scan(&n)
		if n == -1 {
			return
		}
		fmt.Println(calcolaGiornoAnno(giorniDellaSettimana, gds, n))
	}
}

func calcolaGiornoAnno(giorniSettimana [7]string, primoGennaio string, g int) string {
	return giorniSettimana[(g + arrayIndex(giorniSettimana, primoGennaio) - 1) % 7]
}

func arrayIndex(a [7]string, s string) int {
	for i, e := range a {
		if e == s {
			return i
		}
	}
	return -1
}