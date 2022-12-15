package main

import (
    "fmt"
    "io"
)


func main() {
    var list []int
    var num, ricercato int
    fmt.Scan(&ricercato)
    for {
        _, err := fmt.Scan(&num)
        if err == io.EOF {
            break
        }
        list = append(list, num)
    }
    fmt.Println(recursiveCount(ricercato, list))
}

func recursiveCount(el int, list []int) int {
    if len(list) == 0 {
        return 0
    }
    if list[0] == el {
        return recursiveCount(el, list[1:]) + 1
    } else {
        return recursiveCount(el, list[1:])
    }
}