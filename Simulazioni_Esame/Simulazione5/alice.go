package main

import (
	"fmt"
)

type Persona struct {
	nome string
	altezza float64 //in cm
	eta int
}

func main() {
	persone := []Persona{Persona{"Alice", 150.0, 12},Persona{"Pincopanco", 170.0, 15},Persona{"Pancopinco", 170.0, 15}}
	fmt.Println(persone)
	mangiaFungo(&persone[0])
	fmt.Println(entraTana(persone))
}

func mangiaFungo(p *Persona) {
	p.altezza /= 100.0
}

func entraTana(persone []Persona) (entrate []Persona) {
	for _, persona := range persone {
		if persona.altezza < 3.0 {
			entrate = append(entrate, persona)
		}
	}
	return
}