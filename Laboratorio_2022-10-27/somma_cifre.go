package main

import "fmt"

func main() {
  var n int
  fmt.Print("Inserire numero: ")
  fmt.Scan(&n)
  fmt.Println("Somma delle cifre:", sommaCifre(n))
}

func sommaCifre(n int) (s int) {
  for {
    cifra := n % 10
    n /= 10
    s += cifra
    if n == 0 {
      return s
    }
  }
  return -1
}
