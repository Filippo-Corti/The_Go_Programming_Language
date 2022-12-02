package main

import (
	"fmt"
)

func main() {
	var first1, first2 *Node
	
	first2 = AddLast(first2, "paolo")
	first2 = AddLast(first2, "malato")
	first2 = AddLast(first2, "e'")
	first2 = AddLast(first2, "stato")


	first1 = Concat(first1, first2)

	fmt.Printf("Lunghezza = %d\n", Length(first1))
	Print(first1)
}