package main

import "fmt"
//la sequenza DEVE iniziare con 1.
func main() {
	var n, numSequenze, prec int
	for {
		fmt.Scan(&n)
		if n == 2 {
			break
		}
		if n == 0 && prec == 1 {
			numSequenze++
		}
		prec = n
	}
	fmt.Println("Il numero di sottosequenze di zeri è:", numSequenze)
}
