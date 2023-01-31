package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    n, err := strconv.Atoi(os.Args[1])
    if err != nil {
        os.Exit(1)
    }
    righello(n)
}

func righello(n int) {
    if n == 0 {
        return
    }
    righello(n - 1)
    fmt.Println(strings.Repeat("-", n))
    righello(n - 1)
}