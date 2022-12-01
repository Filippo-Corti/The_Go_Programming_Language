package main

import (
    "fmt"
    "strconv"
)

func main() {
    var n int
    var nString, vString string
    var numeriInLettere map[int]string
    numeriInLettere = make(map[int]string)
    fmt.Print("Un numero: ")
    fmt.Scan(&n)
    nString = strconv.Itoa(n)
    //Riempi dizionario
    for _, v := range nString {
        _, ok := numeriInLettere[int(v - '0')]
        if !ok {
            fmt.Printf("Parola per %c? ", v)
            fmt.Scan(&vString)
            numeriInLettere[int(v - '0')] = vString
        }
    }
    //Stampa numero
    fmt.Printf("%s", numeriInLettere[int(nString[0] - '0')])
    for _, v := range nString[1:] {
        fmt.Printf(" - %s", numeriInLettere[int(v - '0')])
    }
    fmt.Println()
}
