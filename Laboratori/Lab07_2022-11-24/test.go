package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	//b := []int{5, 6, 7, 8, 9}

	i := 2
	//j := 4
	w := 10

	a = append(a[:i], append([]int{w}, a[i:]...)...)
	fmt.Println(a)
}