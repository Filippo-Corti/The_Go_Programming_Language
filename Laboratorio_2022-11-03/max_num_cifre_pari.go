package main

import "fmt"

func main() {
	var max int
	var str string
	for i := 0; i < 10; i++ {
		fmt.Print("Numero: ")
		fmt.Scan(&str)
		nCifrePari := contaCifrePari(str)
		if nCifrePari > max {
			max = nCifrePari
		}
	}
	fmt.Println("Massimo:", max)
}

func contaCifrePari(s string) (n int) {
	for i := 0; i < len(s); i++ {
		if (s[i] - '0') % 2 == 0 {
			n++
		} 
	}
	return n
}
