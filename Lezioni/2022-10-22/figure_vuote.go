package main

import "fmt"

func main() {
	triangoloVuoto(100)
}

func quadratoVuoto(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("* ")
	}
	fmt.Println()
	for i := 0; i < n - 2; i++ {
		fmt.Print("* ")
		for j := 0; j < n - 2; j++ {
			fmt.Print("  ")
		}
		fmt.Println("* ")
	}
	for i := 0; i < n; i++ {
		fmt.Print("* ")
	}
	fmt.Println()
}

func triangoloVuoto(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("* ")
	}
	fmt.Println()
	for i := 0; i < n - 2; i++ {
		fmt.Print("* ")
		for j := 0; j < n - i - 3; j++ {
			fmt.Print("  ")
		}
		fmt.Println("* ")
	}
	fmt.Println("* ")
}