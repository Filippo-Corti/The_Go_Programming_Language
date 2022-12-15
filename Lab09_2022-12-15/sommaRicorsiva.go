package main

import (
    "fmt"
    "io"
)

func main() {
    var list []int
    var num int
    for {
        _, err := fmt.Scan(&num)
        if err == io.EOF {
            break
        }
        list = append(list, num)
    }
    fmt.Println(recursiveSum(list))
}

func recursiveSum(list []int) int {
    if len(list) == 0 {
        return 0
    }
    return recursiveSum(list[1:]) + list[0]
}