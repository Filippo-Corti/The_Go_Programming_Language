package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

type Punto struct {
	etichetta string
	x, y float64
}

type Segmento struct {
	start, end Punto
}

func (s Segmento) isSimmetricForYAxis() bool {
	return s.start.y == s.end.y && s.start.x == - s.end.x
}

func (s Segmento) isBelowXAxis() bool {
	return s.start.y < 0 && s.end.y < 0 
}

func StringPunto(p Punto) string {
	return fmt.Sprintf("%s = (%.2f, %.2f)", p.etichetta, p.x, p.y)
}

func StringSegmento(s Segmento) string {
	return fmt.Sprintf("Segmento con estremi %s e %s", StringPunto(s.start), StringPunto(s.end))
}
 
func Distanza(p1, p2 Punto) float64 {
	return math.Sqrt(math.Pow(p1.x - p2.x, 2) + math.Pow(p1.y - p2.y, 2))
}

func main() {
	var valoreSoglia float64
	fmt.Scan(&valoreSoglia)
	punti := readInput(os.Args[1:])
	printSegmentiValidi(punti, valoreSoglia)
}

func printSegmentiValidi(punti []Punto, soglia float64) {
	for i := 0; i < len(punti); i++ {
		for j := i + 1; j < len(punti); j++ {
			printSegmentoValido(Segmento{punti[i], punti[j]}, soglia)
		}
	}
}

func printSegmentoValido(segmento Segmento, soglia float64) {
	if segmento.isSimmetricForYAxis() && segmento.isBelowXAxis() && Distanza(segmento.start, segmento.end) > soglia {
		fmt.Println(StringSegmento(segmento))
	}
}

func readInput(in []string) (punti []Punto) {
	for i := 0; i < len(in); i += 3 {
		x, _ := strconv.ParseFloat(in[i + 1], 64)
		y, _ := strconv.ParseFloat(in[i + 2], 64)
		punti = append(punti, Punto{in[i], x, y})
	}
	return
}