package main

import "fmt"

func main() {
	a := []int{34, 57, 87, 12, 109, 180, 655, 3186}
	fmt.Println(estraiPari(a))
	fmt.Println(rimuoviMultipli(2, a))
}

func estraiPari(in []int) (out []int) {
	for _, v := range in {
		if v % 2 == 0 {
			out = append(out, v)
		}
	}
	return
}

func rimuoviMultipli(m int, in []int) (out []int) {
	for _, v := range in {
		if v % m != 0 {
			out = append(out, v)
		}
	}
	return
}