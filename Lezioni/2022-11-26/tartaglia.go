package main

import (
	"fmt"
	"strings"
)

const SEPARATOR = "     "

func main() {
	var rows int
	fmt.Print("Righe: ")
	fmt.Scan(&rows)
	for row := 1; row <= rows; row++ {
		fmt.Print(strings.Repeat(SEPARATOR, rows-row))
		stampaRigaTartaglia(row)
	}
}

func stampaRigaTartaglia(n int) {
	for i := 0; i < n; i++ {
		v := getTartaglia(n - 1, i)
		fmt.Printf("%5d%s", v, SEPARATOR)
	}
	fmt.Println()
}

func getTartaglia(row, col int) (n int) {
	if col == 0 || col == row {
		return 1
	}
	return getTartaglia(row - 1, col - 1) + getTartaglia(row - 1, col)
}