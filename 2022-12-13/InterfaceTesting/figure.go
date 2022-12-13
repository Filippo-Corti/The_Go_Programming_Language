package main

import (
	"fmt"
	"math"
	"reflect"
)

type FiguraGeometrica interface {
	Area() int
	Perimetro() int
}

type Rettangolo struct {
	l1, l2 int
}

func (r *Rettangolo) Area() int {
	return r.l1*r.l2
}

func (r *Rettangolo) Perimetro() int {
	return (r.l1+r.l2)*2
}

type Triangolo struct {
	base, altezza int
}

func (t *Triangolo) Area() int {
	return (t.base*t.altezza) / 2
}

func (t *Triangolo) Perimetro() int {
	return t.base + 2 * (int(math.Sqrt(float64(t.altezza*t.altezza + (t.base/2)*(t.base/2)))))
}

func main() {
	var figure []FiguraGeometrica
	rett := Rettangolo{l1: 3, l2: 5}
	triang := Triangolo{base: 5, altezza: 9}
	figure = append(figure, FiguraGeometrica(&rett))
	figure = append(figure, FiguraGeometrica(&triang))
	stampaFigure(figure)
}

func stampaFigure(figure []FiguraGeometrica) {
	for i, figura := range figure {
		fmt.Printf("%d: Tipo: %s - Area: %d, Perimetro: %d\n", i, reflect.TypeOf(figura).Elem(), figura.Area(), figura.Perimetro())
	}
}