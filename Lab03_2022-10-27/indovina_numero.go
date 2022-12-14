package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  MAX := 10
  nDaIndovinare := rand.Intn(MAX) + 1
  var tent int
  for i := 0; i < MAX/2; i++ {
    fmt.Print("Tentativo: ")
    fmt.Scan(&tent)
    if tent == nDaIndovinare {
      fmt.Println("Hai indovinato in", i + 1, "tentativi")
      return
    }
    if tent < 1 || tent > MAX {
      fmt.Println("Fuori intervallo!")
      i--;
    }
  }
  fmt.Println("Hai perso, il numero era", nDaIndovinare)
}
