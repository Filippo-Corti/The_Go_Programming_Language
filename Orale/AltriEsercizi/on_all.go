package main

import (
	"fmt"
)

func square(n int) int {
	return n*n
}

func main() {
	list := []int{1,2,3,4,5,6,7}
	fmt.Println(list)
	on_all(list, square)
	fmt.Println(list)
}

func on_all(list []int, function func(int)int) {
	for i := 0; i < len(list); i++ {
		list[i] = function(list[i])
	}
}