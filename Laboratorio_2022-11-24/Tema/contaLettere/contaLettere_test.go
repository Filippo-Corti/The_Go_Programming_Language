/*
	vedi note sul repo
*/

package main

import (
	"fmt"
	//"log"
	//"os/exec"

	//"strings"
	"testing"
	//"bytes"
	//"io"
	//"os"
)

func TestEsistenzaCostante(t *testing.T) {
	fmt.Println("*** questo test non compila se manca la const LEN_ALFABETO:", LEN_ALFABETO)
}

func TestEsistenzaFunzione(t *testing.T) {
	fmt.Println("*** questo test non compila se manca la funzione")
	//aggiornaConteggi("abc", nil)
}

func TestMinimale(t *testing.T) {
	var conteggi [LEN_ALFABETO]int
	aggiornaConteggi("abc", &conteggi)
	if conteggi[0] != 1 {
		t.Fail()
	}
}

var prog = "./contaLettere"

func TestMinimaleMain(t *testing.T) {
	LanciaGenerica(t,
		prog,
		"abc",
		"a 1\nb 1\nc 1\nd 0\ne 0\nf 0\ng 0\nh 0\ni 0\nj 0\nk 0\nl 0\nm 0\nn 0\no 0\np 0\nq 0\nr 0\ns 0\nt 0\nu 0\nv 0\nw 0\nx 0\ny 0\nz 0\n",
		"NIENTE")
}
