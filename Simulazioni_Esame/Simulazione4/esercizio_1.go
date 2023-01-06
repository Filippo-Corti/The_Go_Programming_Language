package main

import (
	"fmt"
)

func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	printDinDonDan(n1, n2)
}

func printDinDonDan(start, end int) {
	for n := start; n <= end; n++ {
		fmt.Print(getDinDonDan(n), " ")
	}
	fmt.Println()
}

func getDinDonDan(n int) (ris string) {
	if n % 3 == 0 {
		ris += "Din"
	}
	if n % 5 == 0 {
		ris += "Don"
	}
	if n % 7 == 0 {
		ris += "Dan"
	}
	if ris != "" {
		return ris
	}
	return fmt.Sprintf("%d", n)
}