package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

type TestCase struct {
	f Funzione
	F Funzione
}

func testThis(t *testing.T, tc TestCase, index int) {
	rand.Seed(time.Now().UnixNano())
	a := rand.Float64() * 2.0 + 1.0
	b := a + rand.Float64() * 2.0 + 1.0
	areaOttenuta := integra(tc.f, a, b, 100000)
	areaAttesa := tc.F(b) - tc.F(a)
	if math.Abs(areaAttesa - areaOttenuta) > 1. {
		t.Error("Not working with", index)
	} 
}

func TestIntegra(t *testing.T) {
	s := []TestCase{
		{math.Sin, func(x float64) float64 {
			return - 1 * math.Cos(x)
		}},
		{math.Exp, math.Exp}}
	for i, tc := range s {
		testThis(t, tc, i)
	}
}