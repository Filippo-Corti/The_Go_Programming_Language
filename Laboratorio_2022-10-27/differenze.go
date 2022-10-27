package main

import "fmt"

func main() {
  var n, nPrec int
  for {
    fmt.Print("Numero: ")
    fmt.Scan(&n)
    if n == 0 {
      break
    }
    fmt.Println("Differenza:", n - nPrec)
    nPrec = n
  }
}
