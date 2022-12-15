package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    numero, _ := strconv.Atoi(os.Args[1])
    fmt.Println(fattoriale(numero))
}

func fattoriale(n int) int {
    if n == 0 {
        return 1
    }
    return fattoriale(n-1) * n
} 