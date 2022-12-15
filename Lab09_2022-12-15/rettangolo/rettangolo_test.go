package main

import (
	"fmt"
	"testing"
)

var prog = "./rettangolo"

func TestEsistenzaMetodo(t *testing.T) {
	var r Rettangolo
	//fmt.Println(r)
	fmt.Println(r.String())
	/*
		if r.String() != "base:0,altezza:0" {
			t.Fail()
		}
	*/
}
