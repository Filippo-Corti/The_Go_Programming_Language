package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
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
    //stampaMappaOrdinata(contatoreVocali, VOCALI)
    stampaMappaOrdinataAlfabeto(contatoreVocali)
}

func stampaMappaOrdinata(m map[rune]int, ordine string) {
    for _, r := range ordine {
        el, ok := m[r]
        if ok {
            fmt.Printf("%c : %d\n", r, el)
        }
    }
}

//Oppure
func stampaMappaOrdinataAlfabeto(m map[rune]int) {
    keys := getKeys(m)
    sort.Strings(keys)
    for _, k := range keys {
        runa := rune(k[0])
        fmt.Printf("%c : %d\n", runa, m[runa])
    }
}   

//Ho provato a usare la MapKeys() di reflect ma non riuscivo a convertire la slice di Value in slice di stringhe/rune
//Ritorno una slice di stringhe, invece che di runes, perché sort.Runes() non esiste
func getKeys(m map[rune]int) (s[]string){
    for k, _ := range m {
        s = append(s, string(k))
    }
    return
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