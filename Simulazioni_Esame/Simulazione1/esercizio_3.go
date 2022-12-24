package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Utenza struct {
	numeroTelefonico string //Non ha molto senso che sia un intero...
	codiceSIM string
}

type RegistroTelefonico map[string][]Utenza


func main() {
	utenze := LeggiUtenze()
	registro := InizializzaRegistro()
	for _, utenza := range utenze {
		registro = AggiungiUtenza(registro, utenza)
	}
	esegui(registro)
}

func esegui(registro RegistroTelefonico) {
	var tot, max, cont int
	var telefoniSenzaCambi []string
	for numeroTelefonico, _ := range registro {
		if strings.HasPrefix(numeroTelefonico, "338") {
			cambi := getNumeroCambiDiSIM(registro, numeroTelefonico)
			if cambi == 0 {
				telefoniSenzaCambi = append(telefoniSenzaCambi, numeroTelefonico)
			} 
			if cambi > max {
				max = cambi
			}
			tot += cambi
			cont++
		}
	}
	fmt.Printf("Media cambio SIM: %.2f\n", float64(tot)/float64(cont))
	fmt.Printf("Numero massimo cambi SIM: %d\n", max)
	fmt.Println("Numeri di telefono senza cambi SIM:")
	for _, tel := range telefoniSenzaCambi {
		fmt.Println(tel)
	}
}

func getNumeroCambiDiSIM(registro RegistroTelefonico, telefono string) (n int) {
	return NumeroUtenze(registro, telefono) - 1
}

func NumeroUtenze(registro RegistroTelefonico, telefono string) (n int) {
	return len(registro[telefono])
}

func Identifica(registro RegistroTelefonico, telefono string) (codiceSIM string) {
	return registro[telefono][len(registro[telefono]) - 1].codiceSIM
}

func AggiungiUtenza(registro RegistroTelefonico, utenza Utenza) (registroAggiornato RegistroTelefonico) {
	registro[utenza.numeroTelefonico] = append(registro[utenza.numeroTelefonico], utenza)
	return registro
}

func InizializzaRegistro() (registro RegistroTelefonico) {
	return make(map[string][]Utenza)
}

func LeggiUtenze() (utenze []Utenza) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ";")
		utenze = append(utenze, Utenza{data[0], data[1]})
	}
	return
}