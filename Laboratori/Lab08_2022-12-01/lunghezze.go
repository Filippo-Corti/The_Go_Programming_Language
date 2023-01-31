package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    var contaLunghezzaParole map[int]int
    contaLunghezzaParole = make(map[int]int)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        aggiornaConteggio(contaLunghezzaParole, line)
    }
    fmt.Println(contaLunghezzaParole)
}

func aggiornaConteggio(m map[int]int, riga string) {
    parole := strings.Split(riga, " ")
    for _, parola := range parole {
        m[len(parola)]++
    }
    delete(m, 0) //Le parole da zero lettere non hanno molto senso quindi le togliamo
}