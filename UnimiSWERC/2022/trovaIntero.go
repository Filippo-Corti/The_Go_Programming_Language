package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	res := (a + b) / 2
	if (a + b) % 2 == 0 {
		fmt.Println(res)
	} else {
		fmt.Println(("IMPOSSIBILE"))
	}
}