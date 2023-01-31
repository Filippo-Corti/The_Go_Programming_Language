package main

import (
    "fmt"
    "os"
    "reflect"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("input errato")
        os.Exit(1)
    }
    fmt.Println("Anagrammi?:", isAnagram(os.Args[1], os.Args[2]))
}

func isAnagram(s1, s2 string) bool {
    return reflect.DeepEqual(mappaCaratteri(s1), mappaCaratteri(s2))
}

func mappaCaratteri(s string) map[rune]int {
    var m map[rune]int
    m = make(map[rune]int)
    for _, char := range s {
        m[char]++
    }
    return m
}