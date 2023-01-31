package main

import "fmt"
import "math/rand"
import "time"

func main() {
  rand.Seed(time.Now().UnixNano())
  s := 0
  for i := 0; i < 10; i++ {
    n := rand.Intn(11)
    s += n
    fmt.Print(n, "\t")
  }
  fmt.Println()
  fmt.Println("Somma:", s)
}
