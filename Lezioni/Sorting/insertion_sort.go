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

func insertionsort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && a[j - 1] > a[j]; j--  {
			a[j - 1], a[j] = a[j], a[j - 1]
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
	insertionsort(array)
	fmt.Println(array)
}
