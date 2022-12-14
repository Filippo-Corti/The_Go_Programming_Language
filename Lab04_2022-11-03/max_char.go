package main

import "fmt"

func main() {
	var max byte = 0
	var c, cb byte
	for i := 0; i < 5; i++ {
		fmt.Print("Carattere: ")
		fmt.Scanf("%c", &c)
		fmt.Scanf("%c", &cb)
		if c > max {
			max = c
		}
	}
	fmt.Println("Maggiore:", string(max))
}