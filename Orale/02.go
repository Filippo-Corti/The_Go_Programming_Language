/*
2. Funzione che data una slice di parole verifica che per ogni parola ci sia almeno una delle possibili rotazioni
    es: "strada" ha come rotazioni "tradas", "radast", "adastr", "dastra", "astrad"
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	parole := []string{"strada", "dastra", "ciao", "aoci"}
	fmt.Println(verificaRotazioni(parole))
}

// Ritorna true se la slice parole contiene almeno una rotazione di parola
// Una stringa è rotazione di parola se:
// 	- è lunga uguale
// 	- è contenuta nella stringa formata da parola + parola
// es: strada -> tradas è lunga come strada ed è contenuta in "stradastrada"
func contieneUnaRotazione(parola string, parole []string) bool {
	for _, parola2 := range parole {
		doppio := parola2 + parola2
		if parola != parola2 && len(parola2) == len(parola) && strings.Contains(doppio, parola) {
			return true
		}
	}
	return false
}

func verificaRotazioni(parole []string) bool {
	for _, parola := range parole {
		/* Soluzione Brutta:
		if !containsAny(parole, generaRotazioni(parola)) {
			return false
		} */
		/* Questa mi permette di non generare tutte le rotazioni ma di controllare direttamente */
		if !contieneUnaRotazione(parola, parole) {
			return false
		}
	}
	return true
}

/* --------------------------------------------- */

func contains(slice []string, str string) bool {
	for _, element := range slice {
		if str == element {
			return true
		}
	}
	return false
}

// Returns true if s1 contains any of the strings in s2
func containsAny(s1, s2 []string) bool {
	for _, str1 := range s1 {
		if contains(s2, str1) {
			return true
		}
	}
	return false
}

// Genera le rotazioni di parola, esclusa parola stessa (!)
func generaRotazioni(parola string) (rotazioni []string) {
	for i := 1; i < len(parola); i++ {
		rotazioni = append(rotazioni, parola[i:] + parola[:i])
	} 
	return rotazioni
}