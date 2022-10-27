package main

import "fmt"

func main() {
  var n int
  fmt.Print("Inserisci numero: ")
  fmt.Scan(&n)
  fmt.Println("Controllo numero primo:", isPrime(n))
}

func isPrime(n int) bool {
  for i := 2; i <= n/2; i++ {
    if n % i == 0 {
      return false
    }
  }
  return true
}
