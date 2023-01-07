package main

import (
	"fmt"
	"unicode"
)

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(verificaPassword(str))
}

func verificaPassword(pw string) string {
	errors := checkErrors(pw) 
	if len(errors) == 0 {
		return "La pw è ben definita!"
	}
	msg := "La pw non è definita correttamente:"
	for _, err := range errors {
		msg += "\n- " + err
	}
	return msg
}

func checkErrors(pw string) (errors []string) {
	if !checkLen(pw) {
		errors = append(errors, "La pw deve avere una lunghezza minima di 6 caratteri ed una lunghezza massima di 15 caratteri")
	}
	var contaMin, contaMai, contaCifre, contaAltro int
	for _, r := range pw {
		if unicode.IsLower(r) {
			contaMin++
		} else if unicode.IsUpper(r) {
			contaMai++
		} else if unicode.IsDigit(r) {
			contaCifre++
		} else {
			contaAltro++
		}
	}
	if contaMin < 2 {
		errors = append(errors, "Almeno 2 caratteri della pw devono rappresentare delle lettere minuscole")
	}
	if contaMai > 3 {
		errors = append(errors, "Al massimo 3 caratteri della pw possono rappresentare delle lettere maiuscole")
	}
	if contaCifre < 2 {
		errors = append(errors, "Almeno 2 caratteri della pw devono rappresentare delle cifre decimali")
	}
	if contaAltro < 2 {
		errors = append(errors, "Almeno 2 caratteri della pw non devono rappresentare lettere o cifre decimali")
	}
	return errors
}

func checkLen(pw string) bool {
	return len(pw) >= 6 && len(pw) <= 15
}