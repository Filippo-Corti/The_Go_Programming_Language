package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func RimuoviCifraPiuAlta(numero string) string {
	return strings.Replace(numero, string(getCifraPiuAlta(numero)), "", 1)
}

func getCifraPiuAlta(numero string) (max rune) {
	for _, runa := range numero {
		if runa > max {
			max = runa
		}
	}
	return
}

func NumeroPiuPiccolo(numero string, n int) string {
	for i := 0; i < n; i++ {
		numero = RimuoviCifraPiuAlta(numero)
	}
	return numero
}

func main() {
	d, _ := strconv.Atoi(os.Args[2])
	fmt.Println(NumeroPiuPiccolo(os.Args[1], d))
}