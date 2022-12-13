package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    var matrice [][]string
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Righe e colonne: ")
    scanner.Scan()
    line := scanner.Text()
    splitLine := strings.Split(line, " ")
    r, err1 := strconv.Atoi(splitLine[0])
    c, err2 := strconv.Atoi(splitLine[1])
    matrice = make([][]string, r)
    if err1 != nil || err2 != nil {
        fmt.Println("Valori non validi")
        os.Exit(1)
    }
    for i := 0; i < r; i++ {
        scanner.Scan()
        line := scanner.Text()
        splitLine := strings.Split(line, " ")
        matrice[i] = make([]string, c)
        matrice[i] = splitLine
    }
    galleggiaAsterischi(&matrice, r, c)
    stampaMatrice(matrice)
}

//Fa un bubblesort per ogni colonna della matrice, invertendo due elementi se quello sopra è un "*" e quello sotto no
func galleggiaAsterischi(m *[][]string, r int, c int) {
    for colonna := 0; colonna < c; colonna++ {
        for i := 0; i < r - 1; i++ {
            for j := 0; j < r - i - 1; j++ {
                if (*m)[j][colonna] != "*" && (*m)[j + 1][colonna] == "*" {
                    (*m)[j][colonna], (*m)[j + 1][colonna] = (*m)[j + 1][colonna], (*m)[j][colonna]
                }
            }
        }
    }
}

func stampaMatrice(m [][]string) {
    for _, r := range m {
        for _, c := range r {
            fmt.Printf("%s ", c)
        }
        fmt.Println()
    }
}