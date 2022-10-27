package main

import "fmt"


func main() {
  fmt.Println("Somma e Andamento Numeri:")
  var n, nPrec, somma int
  for {
    fmt.Scan(&n)
    if n == -1 {
      break
    }
    if n > nPrec {
      fmt.Println("+")
    } else if n < nPrec {
      fmt.Println("-")
    } else {
      fmt.Println("=")
    }
    nPrec = n
    somma += n
  }
  fmt.Println("Somma:", somma)
}
