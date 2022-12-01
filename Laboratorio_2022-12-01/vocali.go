package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

const VOCALI = "AEIOUaeiou"

func main() {
    var contatoreVocali map[rune]int
    contatoreVocali = make(map[rune]int)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        contaVocali(line, contatoreVocali)
    }
    stampaMappa(contatoreVocali)
}

func isVowel(r rune) bool {
    return strings.ContainsRune(VOCALI, r)
}

func contaVocali(s string, vocali map[rune]int) {
    for _, char := range s {
        if isVowel(char) {
            vocali[char]++
        }
    }
}

func stampaMappa(m map[rune]int) {
    for r,i := range m {
        fmt.Printf("%c : %d\n", r, i)
    }
}