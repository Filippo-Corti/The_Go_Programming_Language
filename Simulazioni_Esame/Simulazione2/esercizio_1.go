package main

import (
	"fmt"
	"os"
	"unicode"
)

func TrasformaParola(parola string, posizione int) (parolaTrasformata string) {
	//Unnecessary but better than ifs?
	functions := []func(rune)rune{unicode.ToUpper, unicode.ToLower}
	for i, runa := range parola {
		parolaTrasformata += string(functions[(i + posizione) % 2](runa))
	}
	return
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Print(TrasformaParola(os.Args[i], i - 1), " ")
	}
	fmt.Println()
}