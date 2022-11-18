package main

import "fmt"

func main() {
	slice1 := []string{"ma", "te", "ma", "ti", "ca"}
	slice2 := []string{"ma", "ti", "ta"}
	fmt.Println("Numero di stringhe di slice1 in comune a slice2:", contaStringheInComune(slice1, slice2))
	fmt.Println("Stringhe di slice1 in comune a slice2:", stringheInComune(slice1, slice2))
}

func sliceContains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func contaStringheInComune(s1, s2 []string) (c int) {
	for _, s := range s1 {
		if sliceContains(s2, s) {
			c++
		}
	}
	return
}

func stringheInComune(s1, s2 []string) (s3 []string) {
	for _, s := range s1 {
		if sliceContains(s2, s) { //Se voglio evitare duplicati all'interno di s3 aggiungo "&& !sliceContains(s3, s)"
			s3 = append(s3, s)
		}
	}
	return
}
