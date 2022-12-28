package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Punto struct {
	etichetta string
	ascissa, ordinata float64
}

func NuovoTragitto() (tragitto []Punto) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		dati := strings.Split(line, ";")
		x, _ := strconv.ParseFloat(dati[1], 64)
		y, _ := strconv.ParseFloat(dati[2], 64)
		tragitto = append(tragitto, Punto{dati[0], x ,y})
	}
	return
}

func Distanza(p1, p2 Punto) float64 {
	return math.Sqrt(math.Pow(p1.ascissa - p2.ascissa, 2) + math.Pow(p1.ordinata - p2.ordinata, 2))
}

func Lunghezza(tragitto []Punto) (lun float64) {
	for i := 0; i < len(tragitto) - 1; i++ {
		lun += Distanza(tragitto[i], tragitto[i + 1])
	}
	return
}

func PuntoOltreMeta(tragitto []Punto, meta float64) Punto {
	var i int
	var lun float64
	for i = 0; i < len(tragitto) - 1; i++ {
		lun += Distanza(tragitto[i], tragitto[i + 1])
		if lun > meta {
			break
		}
	}
	return tragitto[i + 1]
}

func String(p Punto) string {
	return fmt.Sprintf("%s = (%.1f, %.1f)", p.etichetta, p.ascissa, p.ordinata)
}

func main() {
	tragitto := NuovoTragitto()
	lunghezza := Lunghezza(tragitto)
	fmt.Printf("Lunghezza percorso: %.3f\n", lunghezza)
	fmt.Printf("Punto oltre metà: %s\n", String(PuntoOltreMeta(tragitto, lunghezza / 2.0)))
}