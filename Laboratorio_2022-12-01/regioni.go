package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Capoluogo struct {
    nome, regione string
}

const POS_NOME = 0
const POS_REGIONE = 2

func main() {
    var regione string
    var capoluoghi map[string]([]Capoluogo)
    capoluoghi = make(map[string]([]Capoluogo))
    //Riempi dizionario capoluoghi
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() //Salta Legenda
    for scanner.Scan(){
        line := scanner.Text()
        campi := strings.Split(line, ",")
        capoluoghi[campi[POS_REGIONE]] = append(capoluoghi[campi[POS_REGIONE]], Capoluogo{nome: campi[POS_NOME], regione: campi[POS_REGIONE]})
    }
    //Interroga dizionario
    for i := 1; i < len(os.Args); i++ {
        regione = os.Args[i]
        listaCapoluoghiInRegione, ok := capoluoghi[regione]
        if !ok {
            fmt.Println("Input", regione, "non valido")
            continue
        }
        fmt.Printf("Capoluoghi della regione %s:\n", regione)
        for _, cap := range listaCapoluoghiInRegione {
            fmt.Printf("%s ", cap.nome)
        }
        fmt.Printf("\n-----------------\n")
    }
}