package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Comando struct {
	direzione string
	passi int
}

func (c Comando) String() string {
	return fmt.Sprintf("%s %d", c.direzione, c.passi)
}

func (c Comando) Inverso() (inverso Comando) {
	switch c.direzione {
	case "NORD":
		inverso.direzione = "SUD"
	case "SUD":
		inverso.direzione = "NORD"
	case "EST":
		inverso.direzione = "OVEST"
	case "OVEST":
		inverso.direzione = "EST"
	}
	inverso.passi = c.passi
	return
}

func main() {
	comandi := LeggiComandi()
	contatori := AnalizzaComandi(comandi)
	stampaRisultati(comandi, contatori)
}

func stampaRisultati(comandi []Comando, contatori map[string]int) {
	fmt.Println("Movimenti totali:")
	for key, cont := range contatori {
		fmt.Printf("%s %d\n", key, cont)
	}
	fmt.Println("Direzione in cui il robot deve compiere il minor numero di passi:")
	fmt.Println(trovaMin(contatori, comandi))
	fmt.Println("Comandi Inversi:")
	for i := len(comandi) - 1; i > 0; i-- {
		fmt.Printf("%v, ", comandi[i].Inverso().String())
	}
	fmt.Printf("%v\n", comandi[0].Inverso().String())
}

func trovaMin(contatori map[string]int, comandi []Comando) string {
	var min int = contatori[comandi[0].direzione]
	var minKey string = comandi[0].direzione
	for key, cont := range contatori {
		if cont < min {
			min = cont
			minKey = key
		}
	}
	return minKey
}

func AnalizzaComandi(comandi []Comando) map[string]int {
	counters := make(map[string]int)
	for _, comando := range comandi {
		counters[comando.direzione] += comando.passi
	}
	return counters
}

func LeggiComandi() (comandi []Comando) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		splittedLine := strings.Fields(line)
		steps, _ := strconv.Atoi(splittedLine[1])
		comandi = append(comandi, Comando{splittedLine[0], steps})
	}
	return
}