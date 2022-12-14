package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    var mappaPosizioniParole map[string]([]int)
    mappaPosizioniParole = make(map[string]([]int))
    letteFinora := 0
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Scrivi parole (ctrl D per terminare)")
    for scanner.Scan() {
        line := scanner.Text() 
        parole := strings.Split(line, " ")
        for i, parola := range parole {
            mappaPosizioniParole[parola] = append(mappaPosizioniParole[parola], letteFinora + i)
        }
        letteFinora += len(parole)
    }
    fmt.Println(mappaPosizioniParole)
}