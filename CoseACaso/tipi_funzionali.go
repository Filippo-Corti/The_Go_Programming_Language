package main

import (
	"fmt"
)

func getFunzioneMoltiplicatore(a int) func(int)int {
	return func(x int) int {
		return x * a
	}
}

func main() {
	multiplier := getFunzioneMoltiplicatore(5);
	fmt.Println(multiplier(7));
}