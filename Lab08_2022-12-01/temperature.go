package main

import (
    "fmt"
    "sort"
)

func main() {
    var n, sommaTemperature, numeroTemperature int
    var listaTemperature []int
    for {
        fmt.Scan(&n)
        if n == 999 {
            break
        }
        sommaTemperature += n
        listaTemperature = append(listaTemperature, n)
    }
    numeroTemperature = len(listaTemperature)
    sort.Ints(listaTemperature)
    fmt.Println("Media:", float64(sommaTemperature) / float64(numeroTemperature))
    fmt.Println("Mediana:", float64((listaTemperature[(numeroTemperature + 1) / 2] + listaTemperature[numeroTemperature / 2 - 1])) / 2.0)
    fmt.Println("Temperature sotto media:", contaTemperatureSottoMedia(listaTemperature, float64(sommaTemperature) / float64(numeroTemperature)))
    if numeroTemperature >= 3 {
        fmt.Println("3 Temperature più basse:", listaTemperature[0:3])
        fmt.Println("3 Temperature più alte:", listaTemperature[numeroTemperature - 3:])
    }
}

func contaTemperatureSottoMedia(lista []int, media float64) (i int) {
    for float64(lista[i]) < media {
        i++
    }
    return 
}