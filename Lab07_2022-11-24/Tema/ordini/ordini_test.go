/*
	vedi note sul repo
*/

package main

import (
	//"fmt"
	//"log"
	//"os/exec"

	//"strings"
	"testing"
	//"bytes"
	//"io"
	//"os"
)

var prog = "./ordini"

func TestMinimalissimo(t *testing.T) {
	LanciaGenerica(t,
		prog,
		"100#1#.1\n",
		"90\n",
		"NIENTE")
}

func TestMinimale(t *testing.T) {
	LanciaGenerica(t,
		prog,
		"100#3#.2\n100#3#.1\n",
		"510\n",
		"NIENTE")
}
