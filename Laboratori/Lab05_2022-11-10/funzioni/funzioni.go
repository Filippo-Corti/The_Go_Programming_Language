package main

import "fmt"
import "unicode"

func main() {
	var s string
	fmt.Scan(&s)
	if hasUpper(s) {
		fmt.Printf("Ha maiuscole\n")
	} else {
		fmt.Printf("Non ha maiuscole\n")
	}
	fmt.Printf("Prima maiuscola in posizione: %d\n", firstUpper(s))
	fmt.Printf("Ultima maiuscola in posizione: %d\n", lastUpper(s))
	d, l, o := countDigitsLettersOthers(s)
	fmt.Printf("Cifre: %d - Lettere: %d - Altro: %d\n", d, l, o)
	if isPalindrome(s) {
		fmt.Printf("E' palindroma\n")
	} else {
		fmt.Printf("Non è palindroma\n")
	}
}

/*
hasUpper(s string) bool. 
La funzione riceve una stringa come parametro e restituisce true se la stringa
contiene almeno una lettera latina maiuscola (tra 'A' e 'Z'), false altrimenti.
*/
func hasUpper(s string) bool {
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			return true
		}
	}
	return false
}

/*
firstUpper(s string) int. 
La funzione riceve una stringa come parametro e restituisce la posizione 
della prima lettera latina maiuscola (tra 'A' e 'Z'), se la stringa 
ne contiene almeno una, -1 altrimenti.
*/
func firstUpper(s string) int {
	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			return i
		}
	}
	return -1
}

/*
Scrivere una funzione lastUpper(s string) int. 
La funzione riceve una stringa come parametro e 
restituisce la posizione dell'ultima lettera latina 
maiuscola (tra 'A' e 'Z'), se la stringa ne contiene 
almeno una, -1 altrimenti.
*/
func lastUpper(s string) int {
	len := len(s)
	for i := len - 1; i >= 0; i-- {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return i
		}
	}
	return -1
}

/*
Scrivere una funzione countDigitsLettersOthers(s string) int int int. 
La funzione riceve una stringa come parametro, conta quante cifre, 
quante lettere e quanti caratteri che non sono né cifre né lettere 
contiene e restituisce i tre risultati in questo ordine.
*/
func countDigitsLettersOthers(s string) (d int, l int, o int) {
	for _, r := range s {
		switch {
		case unicode.IsDigit(r):
			d++
		case unicode.IsLetter(r):
			l++
		default:
			o++
		} 
	}
	return
}

/*
Scrivere una funzione isPalindrome(s string) bool. 
La funzione riceve una stringa come parametro e restituisce 
true se la stringa è palindroma, e false altrimenti. Una parola è 
palindroma se può essere letta  sia da sinistra a destra che da destra a sinistra. 
Ad esempio "osso" e "ingegni" sono palindrome, ma anche "" e "t".
*/
func isPalindrome(s string) bool {
	for i := 0; i < len(s) / 2; i++ {
		if s[i] != s[len(s) - 1 - i]{
			return false
		}
	}
	return true
}


