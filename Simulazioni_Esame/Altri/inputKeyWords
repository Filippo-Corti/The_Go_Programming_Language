package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers := parseToFloats(os.Args[1:])
	compareSlice(numbers[1:], numbers[0])
}

func compareSlice(slice []float64, epsilon float64) {
	for i := 0; i < len(slice) - 1; i++ {
		switch {
		case slice[i + 1] - slice[i] > epsilon:
			fmt.Print(">")
		case slice[i + 1] - slice[i] < - epsilon:
			fmt.Print("<")
		default:
			fmt.Print("=")
		}
	}
	fmt.Println()
}

func parseToFloats(list []string) (floats []float64) {
	for _, el := range list {
		value, _ := strconv.ParseFloat(el, 64)
		floats = append(floats, value)
	}
	return
}

