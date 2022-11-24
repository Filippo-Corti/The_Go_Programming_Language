package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	dati := strings.Split(input, " ")
	fmt.Println(valutaPolinomio(dati[:len(dati) - 1], dati[len(dati) - 1]))
}

func valutaPolinomio(polinomio []string, x string) (t float64) {
	xInt, _ := strconv.ParseFloat(x, 64)
	for i, s := range polinomio {
		value, _ := strconv.ParseFloat(s, 64)
		t += value * math.Pow(xInt, float64(i))
	}
	return
}