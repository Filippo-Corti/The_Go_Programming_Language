package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"strconv"
)

type Punto struct {
	etichetta string
	ascissa, ordinata float64
}

func (p Punto) String() string {
	return fmt.Sprintf("%s = (%.1f, %.1f)", p.etichetta, p.ascissa, p.ordinata)
}

func (p Punto) getQuadrante() int {
	if p.ascissa >= 0 {
		if p.ordinata >= 0 {
			return 1
		} else {
			return 4
		}
	} else {
		if p.ordinata >= 0 {
			return 2
		} else {
			return 3
		}
	}
}

type Triangolo struct {
	a, b, c Punto
}

func (t Triangolo) String() string {
	return fmt.Sprintf("Triangolo rettangolo con vertici %v, %v, %v, ed area %.1f.", t.a, t.b, t.c, t.Area())
}

func (t Triangolo) Area() float64 {
	//Usando la Formula di Erone
	a := Distanza(t.a, t.b)
	b := Distanza(t.b, t.c)
	c := Distanza(t.a, t.c)
	p := (a + b + c) / 2
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

//La funzione considera impropriamente triangoli rettangolo solo quelli con cateti paralleli agli assi
func (t Triangolo) isRettangolo() bool {
	var contaParalleli int
	if isParalleloAdUnAsse(t.a, t.b) {
		contaParalleli++
	}
	if isParalleloAdUnAsse(t.a, t.c) {
		contaParalleli++
	}
	if isParalleloAdUnAsse(t.b, t.c) {
		contaParalleli++
	}
	return contaParalleli == 2
}

func isParalleloAdUnAsse(p1, p2 Punto) bool {
	return p1.ascissa == p2.ascissa || p1.ordinata == p2.ordinata 
}

//Per Distribuiti si intende che appartengano almeno a 2 quadranti diversi
func (t Triangolo) arePuntiDistribuiti() bool {
	q1 := t.a.getQuadrante()
	q2 := t.b.getQuadrante()
	q3 := t.c.getQuadrante()
	return !(q1 == q2 && q1 == q3)
}
  
func main() {
	punti := LeggiPunti()
	triangoli := GeneraTriangoli(punti)
	ok, t := cercaTriangoloRettangoloMinore(triangoli)
	if ok {
		fmt.Println(t)
	}
}

func cercaTriangoloRettangoloMinore(triangoli []Triangolo) (bool, Triangolo) {
	var found bool
	var minT Triangolo = triangoli[0]
	var minA float64 = triangoli[0].Area()
	for _, triangolo := range triangoli {
		if triangolo.isRettangolo() && triangolo.arePuntiDistribuiti() {
			found = true
			if triangolo.Area() < minA {
				minA = triangolo.Area()
				minT = triangolo
			}
		}
	}
	return found, minT
}

func Distanza(p1, p2 Punto) float64 {
	return math.Sqrt(math.Pow(p1.ascissa - p2.ascissa, 2) + math.Pow(p1.ordinata - p2.ordinata, 2))
}

func GeneraTriangoli(punti []Punto) (triangoli []Triangolo) {
	for i := 0; i < len(punti); i++ {
		for j := i + 1; j < len(punti); j++ {
			for k := j + 1; k < len(punti); k++ {
				triangoli = append(triangoli, Triangolo{punti[i], punti[j], punti[k]})
			}
		}
	}
	return
}

func LeggiPunti() (punti []Punto) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		splittedLine := strings.Split(line, ";")
		x, _ := strconv.ParseFloat(splittedLine[1], 64)
		y, _ := strconv.ParseFloat(splittedLine[2], 64)
		punti = append(punti, Punto{splittedLine[0], x, y})
	}
	return
}