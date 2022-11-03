package main

import "fmt"

func main() {
	var car byte
	fmt.Print("Inserire un Byte: ")
	fmt.Scanf("%c", &car)
	fmt.Println("Byte inserito:", string(car - 1))
	fmt.Println("Byte inserito:", string(car))
	fmt.Println("Byte inserito:", string(car + 1))

	if car >= 'A' && car <= 'L' {
		fmt.Println("A-L")
	} else {
		fmt.Println("altro")
	}

	var str string
	fmt.Print("Una stringa: ")
	fmt.Scan(&str)

	for _, runa := range str {
		fmt.Println(string(runa))
	}


	
}