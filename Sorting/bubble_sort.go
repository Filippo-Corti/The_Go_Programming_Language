package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const MAX = 100

func creaArrayRandom(n int) (a []int) {
	for i := 0; i < n; i++ {
		a = append(a, rand.Intn(MAX))
	}
	return a;
}

//anyChange è una sentinella: se non effettuo alcuno scambio durante una scansione è perché l'array è già in ordine
func bubblesort(a []int) {
	anyChange := true
	for i := 0; i < len(a) && anyChange; i++ {
		anyChange = false
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
				anyChange = true
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		os.Exit(1);
	}
	array := creaArrayRandom(n)
	fmt.Println(array)
	bubblesort(array)
	fmt.Println(array)
}
