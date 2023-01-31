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
    fmt.Println(recursiveMax(list))
}

func recursiveMax(list []int) int {
    if len(list) == 1{
        return list[0]
    }
    return greater(list[0], recursiveMax(list[1:]))
}

func greater(m, n int) int {
    if m > n {
        return m
    }
    return n
}