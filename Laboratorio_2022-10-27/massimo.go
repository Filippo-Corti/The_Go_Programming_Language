package main

import "fmt"
import "math/rand"
import "time"

func main() {
  rand.Seed(time.Now().UnixNano())
  max := 0
  for i := 0; i < 10; i++ {
    n := rand.Intn(21) + 10 // Da 0 a 20 -> da 10 a 30
    if n > max {
      max = n
    }
    fmt.Print(n, "\t")
  }
  fmt.Println()
  fmt.Println("Max:", max)

}
