package main

import (
	"errors"
	"fmt"
)


func main() {
	fmt.Println(funzioneCheRitornaErrore("ciao"))
	fmt.Println(funzioneCheRitornaErrore("non ciao"))
}

func funzioneCheRitornaErrore(str string) error {
	if str == "ciao" {
		return nil
	} else {
		return errors.New("Valore non valido")
	}
}