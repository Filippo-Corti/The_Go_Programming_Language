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

func selectionsort(a []int) {
	for i := 0; i < len(a); i++ {
		minIndex := i
		for j := i + 1; j < len(a); j++ {
			if (a[j] < a[minIndex]) {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
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
	selectionsort(array)
	fmt.Println(array)
}
