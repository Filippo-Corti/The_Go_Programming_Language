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

func TestMinimalePositivo(t *testing.T) {
	if !isSquare(1) {
		t.Fail()
	}
}
func TestMinimaleNegativo(t *testing.T) {
	if isSquare(3) {
		t.Fail()
	}
}

var prog = "./quadrati"

func TestMinimaleMain(t *testing.T) {
	LanciaGenerica(t,
		prog,
		"NIENTE",
		"1\n9\n81\n",
		"1", "5", "9", "81", "7")
}
