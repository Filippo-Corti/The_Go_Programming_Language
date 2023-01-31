package main

import "fmt"

func main() {
  var capitale, interesse, obiettivo float64
  fmt.Print("Inserire dati (capitale interesse(%) obiettivo): ")
  fmt.Scan(&capitale, &interesse, &obiettivo)
  fmt.Println("Anni impiegati a raggiungere l'obiettivo: ", calcolaAnni(capitale, interesse, obiettivo))
}

func calcolaAnni(capitale, interesse, obiettivo float64) (anni int) {
  for {
    capitale += (capitale * (interesse / 100))
    anni++
    if capitale >= obiettivo {
      return anni
    }
    if anni >= 1000 {
      return -1
    }
  }
  return -1
}
