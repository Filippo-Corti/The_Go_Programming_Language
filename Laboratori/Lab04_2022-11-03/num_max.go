package main

import "fmt"

func main() {
	var max, n, c int
	for i := 0; i < 10; i++ {
		fmt.Print("Numero: ")
		fmt.Scan(&n)
		if n > max {
			max = n
			c = 0
		} 
		if n == max {
			c++
		}
	}
	fmt.Println("Massimo:", max, "(Ripetuto", c, "volte)")
}