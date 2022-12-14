package main

import "fmt"
import "unicode"

func main() {
	var s string
	var flagMinu, flagMaiu bool
	flagMinu = true //Tutte minuscole
	flagMaiu = true //Tutte maiuscole
	fmt.Print("Stringa: ")
	fmt.Scan(&s)

	for _, r := range s {
		if unicode.IsLower(r) {
			//R è minuscola
			flagMaiu = false
		}
		if unicode.IsUpper(r) {
			//R è maiuscola
			flagMinu = false
		}
	}
	//Controllo dei flag:
	if flagMinu {
		fmt.Println("Solo minuscole")
	}
	if flagMaiu {
		fmt.Println("Solo maiuscole")
	}
	if !(flagMinu) && !(flagMaiu) {
		fmt.Println("Sia minuscole che maiuscole")
	}


}