// Trasformare una stringa in una stringa palindroma aggiungendo il minor numero di caratteri possibili
package main

import (
	"fmt"
	"os"
)

func isPalindroma(stringa string) bool {
	for i := 0; i < len(stringa) / 2; i++ {
		if stringa[i] != stringa[len(stringa) - i - 1] {
			return false
		}
	}
	return true
}

func palindromizza(stringa string) string {
	if isPalindroma(stringa) {
		return stringa
	}
	for i := 1; i < len(stringa); i++ {
		addedToTail := addToTail(stringa, i)
		addedToHead := addToHead(stringa, i)
		if isPalindroma(addedToHead) {
			return addedToHead
		}
		if isPalindroma(addedToTail) {
			return addedToTail
		}
	}
	return ""
}

func main(){
	fmt.Println(palindromizza(os.Args[1]))
}

//Returns a string composed by adding the first n bytes of str to str itself
//The characters are added in the order shown in the example:
// paolo, 2 -> paoloap (not paolopa)
func addToTail(str string, n int) string {
	for i := 0; i < n; i++ {
		str = str + string(str[n - i - 1])
	}
	return str
}
//Returns a string composed by adding str to the last n bytes of str itself
//The characters are added in the order shown in the example:
// paolo, 2 -> olpaolo (not lopaolo)
func addToHead(str string, n int) string {
	for i := 0; i < n; i++ {
		str = string(str[len(str) - n + i]) + str
	}
	return str
}