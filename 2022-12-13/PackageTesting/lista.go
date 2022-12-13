package main

import (
	"fmt"
	"github.com/Filippo-0/Liste"
)

func main() {
	fmt.Println("Ciao")
	var lista1 *(liste.Node)
	lista1 = lista1.AddFront("Pippo")
	lista1.Print()
}