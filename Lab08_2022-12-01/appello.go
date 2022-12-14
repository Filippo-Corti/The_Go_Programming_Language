package main

import (
    "fmt"
    "os"
    "sort"
)

func main() {
    sort.Strings(os.Args[1:])
    fmt.Println(os.Args[1:])
}