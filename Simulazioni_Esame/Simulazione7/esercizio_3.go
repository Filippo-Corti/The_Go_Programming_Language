package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Studente struct {
	id string
	nome, cognome string
}

func (s Studente) String() string {
	return s.nome + " " + s.cognome
}

func creaStudente(id string) (studente Studente) {
	campiStudente := strings.Split(id, ".")
	nome := campiStudente[0]
	cognome := removeDigits(campiStudente[1])
	studente.id, studente.nome, studente.cognome = id, nome, cognome
	return
}

func parseStudenti(s string) (listaPrenotazione []Studente) {
	studenti := strings.Fields(s)
	for _, studente := range studenti {
		listaPrenotazione = append(listaPrenotazione, creaStudente(studente))
	}
	return
}

func printStudenti(lista []Studente) {
	fmt.Printf("studenti: ")
	for i := 0; i < len(lista) - 1; i++ {
		fmt.Printf("%v,", lista[i])
	}
	fmt.Printf("%v\n", lista[len(lista) - 1])
}

func main() {
	frequenze := make(map[Studente]int)
	capienza, _ := strconv.Atoi(os.Args[1])
	scanner := bufio.NewScanner(os.Stdin) 
	for scanner.Scan() {
		studentiString := scanner.Text()
		studenti := parseStudenti(studentiString)
		studentiAmmessiInClasse := findStudentsAllowedInClass(studenti, capienza, frequenze)
		aggiornaMappaFrequenze(frequenze, studentiAmmessiInClasse)
		printStudenti(studentiAmmessiInClasse)
	}
}

func findStudentsAllowedInClass(studenti []Studente, capienza int, frequenze map[Studente]int) (ammessi []Studente) {
	//I first fill the list
	if len(studenti) <= capienza {
		return studenti
	}
	ammessi = studenti[len(studenti) - capienza :] 
	//I then check if I need to make any changes
	for i := len(studenti) - capienza - 1; i >= 0; i-- {
		ammessi = updateAllowedList(studenti[i], ammessi, frequenze)
	}
	return
}

func updateAllowedList(studente Studente, ammessi []Studente, frequenze map[Studente]int) []Studente {
	for i, studenteAmmesso := range ammessi {
		if frequenze[studente] < frequenze[studenteAmmesso] {
			return append([]Studente{studente}, append(ammessi[:i], ammessi[i + 1:]...)...)
		}
	}
	//Student is not allowed
	return ammessi 
}

func aggiornaMappaFrequenze(frequenze map[Studente]int, studentiAmmessi []Studente) {
	for _, studente := range studentiAmmessi {
		frequenze[studente]++
	}
}

func removeDigits(str string) (newStr string) {
	for _, char := range str {
		if !unicode.IsDigit(char) {
			newStr += string(char)
		}
	}
	return
}