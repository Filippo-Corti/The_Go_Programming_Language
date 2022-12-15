package main

import (
	//"fmt"
	"testing"
)

var prog = "./max"

func TestRico1(t *testing.T) {
	l := []int{3, 5, 6, 8, 9, 9, 10000, 5, 6, 6, 6, 5, 5}
	if recursiveMax(l) != 10000 {
		t.Fail()
	}
}

func TestRico2(t *testing.T) {
	l := []int{3324, 5, 6, 86788, 9, 9, 10000, 5, 6, 4242346, -66666, -5, 5}
	if recursiveMax(l) != 4242346 {
		t.Fail()
	}
}
