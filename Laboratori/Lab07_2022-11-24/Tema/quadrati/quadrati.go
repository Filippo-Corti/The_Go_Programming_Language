package main

import (
	"fmt"
	"os"
	"math"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Numero argomenti insufficiente")
		os.Exit(1)
	}
	for i := 1; i < len(os.Args); i++ {
		v, _ := strconv.ParseFloat(os.Args[i], 64)
		if isSquare(v) {
			fmt.Println(v)
		}
	}
}

func isSquare(n float64) bool {
	return math.Sqrt(n) == math.Trunc(math.Sqrt(n)) 
}