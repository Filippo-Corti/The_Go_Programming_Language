package main

import "fmt"

func main() {
	var prec rune
	var s string
	fmt.Print("Stringa: ")
	fmt.Scan(&s)
	for i, r := range s {
		if i != 0 {
			if r > prec {
				fmt.Print("+")
			} else if r < prec {
				fmt.Print("-")
			} else {
				fmt.Print("=")
			}
		}
		fmt.Print(string(r))
		prec = r
	}
	fmt.Println()

}