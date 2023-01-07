package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	N, _ := strconv.Atoi(os.Args[1])
	k, _ := strconv.Atoi(os.Args[2])
	numeri := FiltraNumeri(GeneraNumeri(N, k))
	for _, num := range numeri {
		fmt.Println(num)
	}
}

func GeneraNumeri(N, k int) (numeri []int) {
	NString := strconv.Itoa(N)
	for i := 0; i <= len(NString) - k; i++ {
		numeroString := NString[:i] + NString[i + k:]
		numeroInt, _ := strconv.Atoi(numeroString)
		numeri = append(numeri, numeroInt)
	}
	return
}

func FiltraNumeri(s1 []int) (s2 []int) {
	for _, n := range s1 {
		if n % 2 != 0 {
			s2 = append(s2, n)
		}
	}
	return
}