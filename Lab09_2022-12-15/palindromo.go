package main

import (
    "fmt"
    "os"
)

func main() {
    if isPalindrome(os.Args[1]) {
        fmt.Printf("%s è palindroma\n", os.Args[1])
    } else {
        fmt.Printf("%s non è palindroma\n", os.Args[1])
    }
}

func isPalindrome(s string) bool {
    if len(s) <= 2 {
        return true
    }
    return isPalindrome(s[1:len(s)-1]) && s[0] == s[len(s)-1]
}