package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Inserire DUE nomi di file")
        os.Exit(1)
    }
    file1, _ := os.ReadFile(os.Args[1])
    file2, _ := os.ReadFile(os.Args[2])
    mescola(strings.Split(string(file1), " "), strings.Split(string(file2), " "))
}

func mescola(slice1, slice2 []string) {
    len1 := len(slice1)
    len2 := len(slice2)
    for i := 0; i < max(len1, len2); i++ {
        if i < len1 {
            fmt.Println(slice1[i])
        }
        if i < len2 {
            fmt.Println(slice2[i])
        }
    }
}

func max(n1, n2 int) int {
    if n1 > n2 {
        return n1
    }
    return n2
}