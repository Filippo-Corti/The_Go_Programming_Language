package main

import (
  "fmt"
  "math"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  var n int
  vicino := 1000
  TARGET := rand.Intn(21)
  for i := 0; i < 5; i++ {
    fmt.Print("Numero ", i+1 , ": ")
    fmt.Scan(&n)
    if math.Abs(float64(TARGET - n)) < math.Abs(float64(TARGET - vicino)) {
      vicino = n
    }
  }
  fmt.Println("Il numero da indovinare era", TARGET)
  fmt.Println("Il numero più vicino ad esso era", vicino)
}
