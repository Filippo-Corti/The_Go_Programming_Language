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
    var capoluoghi map[string]Capoluogo
    capoluoghi = make(map[string]Capoluogo)
    //Riempi dizionario capoluoghi
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() //Salta Legenda
    for {
        scanner.Scan() 
        line := scanner.Text()
        campi := strings.Split(line, ",")
        if len(campi) < 2 {
            break
        }
        superficieInt, _ := strconv.Atoi(campi[POS_SUPERFICIE])
        capoluoghi[campi[POS_SIGLA]] = Capoluogo{nome: campi[POS_NOME], sigla: campi[POS_SIGLA], regione: campi[POS_REGIONE], superficie: superficieInt}
    }
    //Interroga dizionario
    reader := bufio.NewReader(os.Stdin)
    var sigla string
    fmt.Print("Sigla di Capoluogo: ")
    n, _ := fmt.Fscanf(reader, "%s", &sigla)
    fmt.Println(n, sigla)
    for scanner.Scan() {
        fmt.Print("Sigla di Capoluogo: ")
        sigla = scanner.Text()
        if sigla == "stop"{
            break
        }
        capoluogo, ok := capoluoghi[sigla]
        if !ok {
            fmt.Println("Input non valido")
            continue
        }
        fmt.Printf("Dati per %s:\n", capoluogo.nome)
        fmt.Printf("Sigla: %s\nRegione: %s\nSuperficie: %d mq\n", capoluogo.sigla, capoluogo.regione, capoluogo.superficie)
    }
}