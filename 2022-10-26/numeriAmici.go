package main

import (
  "fmt"
  "math"
)

func main() {
  var x, n int
  fmt.Print("Inserire numero x: ")
  fmt.Scan(&x)
  fmt.Print("Inserire numero massimo n: ")
  fmt.Scan(&n)
  fmt.Println("Numeri amici di x con x <= n:")
  findAndPrintFriends(x, n)
}

func findAndPrintFriends(x, n int) {
  for i := 0; i <= n; i++ {
    if areTheyFriends(x, i) {
      fmt.Print(i, "\t")
    }
  }
  fmt.Println()
}

func areTheyFriends(x,y int) bool {
  return x == getSumOfDivisors(y) && y == getSumOfDivisors(x)
}

func getSumOfDivisors(n int) (s int) {
  s += 1
  for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
    if n % i == 0 {
      s += i
      if i != n / i {
        s += (n / i)
      }
    }
  }
  return s
}
