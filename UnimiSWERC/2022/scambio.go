package main

import (
	"fmt"
)

func main() {
	ints := []int{}
	for {
		var n int
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		ints = append(ints, n)
	}

	scambiato := false

	for i := 0; i < len(ints); i++ {
		if ints[i] != i+1 {
			if scambiato {
				fmt.Println("NO")
				return
			}
			other := ints[ints[i] - 1];
			if other == i+1 {
				scambiato = true
				ints[i], ints[ints[i] - 1] = ints[ints[i] - 1], ints[i]
			} else {
				fmt.Println("NO")
				return
			}
		}
	}

	fmt.Println("YES")

}
