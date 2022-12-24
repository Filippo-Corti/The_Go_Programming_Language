package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Count(os.Args[1], os.Args[2]))
}