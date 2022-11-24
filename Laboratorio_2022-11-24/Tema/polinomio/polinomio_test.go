/*
	vedi note sul repo
*/

package main

import (
	"testing"
)

var prog = "./polinomio"

func TestMinimale(t *testing.T) {
	LanciaGenerica(t,
		prog,
		"3 2 0 7 5\n",
		"888\n",
		"NIENTE")
}

func TestGrosso(t *testing.T) {
	LanciaGenerica(t,
		prog,
		"3 2 0 7 8 9 6\n", // 3*6^0+2*6^1+0*6^2+7*6^3+8*6^4+9*6^5=81879
		"81879\n",
		"NIENTE")
}
