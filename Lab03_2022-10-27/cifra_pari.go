package main

import "fmt"

func main() {
  var n int
  fmt.Print("Inserisci numero: ")
  fmt.Scan(&n)
  fmt.Println("Posizione della prima cifra pari:", trovaPosPrimaCifraPari(n))
}

func trovaPosPrimaCifraPari(n int) (c int) {
  for {
    c++
    cifra := n % 10
    n /= 10
    if cifra % 2 == 0 {
      return c
    }
    if n == 0 {
      break
    }
  }
  return -1
}
