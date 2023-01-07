package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	fmt.Println(mcm(a, b))
}

//mcm è a * b fratto il loro MCD
func mcm(a, b int) int {
	return a * b / mcd(a, b)
}

//MCD con Euclide
func mcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}