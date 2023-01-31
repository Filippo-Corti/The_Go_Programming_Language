package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

const MAX_CARICO = 15

type Carrello struct {
    pos int
    carico int
}

func (c Carrello) String() string {
    return fmt.Sprintf("posizione %d, carico %d", c.pos, c.carico)
}


func aggiornaStato(c *Carrello, posizione, carico int) bool {
    if posizione < 0 || carico < 0 || carico > MAX_CARICO {
        return false
    }
    c.pos, c.carico = posizione, carico     
    return true  
}

func main() {
    var carrello Carrello
    if len(os.Args) < 2 {
        fmt.Println("manca nome file")
        os.Exit(1)
    }
    path := getPercorso(os.Args[1])
    stampaPercorso(path)
    scarichi, pesoMax, mappaOggetti := scorriCarrello(&carrello, path)
    stampaPercorso(path)
    fmt.Println("carrello:", carrello.String())
    fmt.Println("n viaggi:", scarichi)
    fmt.Println("peso max:", pesoMax)
    stampaMappaInOrdine(mappaOggetti)
}

func stampaMappaInOrdine(m map[int]int) {
    fmt.Println("oggetti trovati:")
    chiavi := getChiaviOrdinate(m)
    for _, i := range chiavi {
        fmt.Printf("%d ogg. di peso %d\n", m[i], i)
    }
}

func getChiaviOrdinate(m map[int]int) (slice []int) {
    for k, _ := range m {
        slice = append(slice, k)
    }
    sort.Ints(slice)
    return
}

func scorriCarrello(carrello *Carrello, path []int) (contaScarichi, pesoMax int, mappaOggetti map[int]int) {
    mappaOggetti = make(map[int]int)
    contaScarichi++
    for pos := 0; pos < len(path); {
        //Verifico che il carrello possa caricare il nuovo carico
        if aggiornaStato(carrello, pos, carrello.carico + path[pos]) {
            //Carrello carica il carico in posizione pos e avanza
            if path[pos] > pesoMax {
                pesoMax = carrello.carico
            }
            if path[pos] != 0 {
                mappaOggetti[path[pos]]++
            }
            path[pos] = 0
            pos++
        } else {
            //Carrello è pieno: stampa, vado a scaricare e riparto da dove ero rimasto
            fmt.Println("carrello:", carrello.String())
            scaricaCarrello(carrello, path, pos)
            stampaPercorso(path)
            contaScarichi++
        }
    }
    fmt.Println("carrello:", carrello.String())
    scaricaCarrello(carrello, path, 0)
    return
}

func scaricaCarrello(carrello *Carrello, path []int, posCorrente int) {
    path[0] += carrello.carico
    aggiornaStato(carrello, posCorrente, 0) //Aggiorno il carico
}


func getPercorso(nomeFile string) (path []int) {
    file, _ := os.Open(nomeFile)
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    pathString := scanner.Text()
    pathSliceString := strings.Split(pathString, "|")
    for i, _ := range pathSliceString {
        value, _ := strconv.Atoi(pathSliceString[i])
        path = append(path, value)
    }
    return path[1:len(path) - 1]
}

func stampaPercorso(path []int) {
    fmt.Print("|")
    for _, v := range path {
        if v == 0 {
            fmt.Printf(" |")
        } else {
            fmt.Printf("%d|", v)
        }
    }
    fmt.Println()
}