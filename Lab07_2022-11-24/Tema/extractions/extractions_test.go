/*
	vedi note sul repo
*/

package main

import (
	"fmt"
	"reflect"
	"testing"
)

var prog = "./extractions"

func TestPari(t *testing.T) {
	iniziale := []int{1, 4, 7, 8}
	estrazioneExpected := []int{4, 8}
	estrazione := estraiPari(iniziale)
	if !reflect.DeepEqual(estrazione, estrazioneExpected) {
		t.Fail()
	}
}

func TestRimuove(t *testing.T) {
	iniziale := []int{1, 4, 10, 7, 8, 100}
	estrazioneExpected := []int{1, 4, 7, 8}
	estrazione := rimuoviMultipli(10, iniziale)
	fmt.Println(iniziale, estrazioneExpected, estrazione)
	if !reflect.DeepEqual(estrazione, estrazioneExpected) {
		t.Fail()
	}
}

// func TestGrosso(t *testing.T) {
// 	LanciaGenerica(t,
// 		prog,
// 		"3 2 0 7 8 9 6\n", // 3*6^0+2*6^1+0*6^2+7*6^3+8*6^4+9*6^5=81879
// 		"81879\n",
// 		"NIENTE")
// }
