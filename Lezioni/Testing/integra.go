package main

import (
	"fmt"
	"math"
)

type Funzione func(float64) float64 

func integra(f Funzione, a,b float64, n int) (s float64) {
	delta := (b - a)/float64(n)
	for i := 0; i < n; i++ {
		sx := a + delta*float64(i)
		dx := sx + delta
		s += (f(sx) + f(dx))*delta/2
	}
	return s
}

func main() {
	fmt.Println(integra(func(x float64) float64 {
		return 1 - x*x
	}, -1., 1., 100))
	fmt.Println(integra(math.Sin, 0, math.Pi/2., 100))
}