package main

import (
	"fmt"
	"regexp"
)

func main() {
	r, _ := regexp.Compile("[0-9]+")
	fmt.Println(r.FindAllString("5 + 2 + 3 - 667", -1))
}