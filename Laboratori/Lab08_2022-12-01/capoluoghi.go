package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Capoluogo struct {
    nome, sigla, regione string
    superficie int
}

const POS_NOME = 0
const POS_SIGLA = 1
const POS_REGIONE = 2
const POS_SUPERFICIE = 3

func main() {
    var sigla string
    var capoluoghi map[string]Capoluogo
    capoluoghi = make(map[string]Capoluogo)
    //Riempi dizionario capoluoghi
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() //Salta Legenda
    for scanner.Scan() {
        line := scanner.Text()
        campi := strings.Split(line, ",")
        superficieInt, _ := strconv.Atoi(campi[POS_SUPERFICIE])
        capoluoghi[campi[POS_SIGLA]] = Capoluogo{nome: campi[POS_NOME], sigla: campi[POS_SIGLA], regione: campi[POS_REGIONE], superficie: superficieInt}
    }
    //Interroga dizionario
    for i := 1; i < len(os.Args); i++ {
        sigla = os.Args[i]
        capoluogo, ok := capoluoghi[sigla]
        if !ok {
            fmt.Println("Input", sigla, "non valido")
            continue
        }
        fmt.Printf("Dati per %s:\n", capoluogo.nome)
        fmt.Printf("Sigla: %s\nRegione: %s\nSuperficie: %d mq\n", capoluogo.sigla, capoluogo.regione, capoluogo.superficie)
        fmt.Printf("-----------------\n")
    }
}