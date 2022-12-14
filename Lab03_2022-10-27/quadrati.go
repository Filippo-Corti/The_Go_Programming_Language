package main

import "fmt"

func main() {
  var n int
  fmt.Print("Numero: ")
  fmt.Scan(&n)
  fmt.Println("Quadrato più grande <= n: ", trovaQuadratoMinore(n))
}

func trovaQuadratoMinore(n int) int {
  for i := 1; i < n; i++ {
    q := i * i;
    if q > n {
      return (i - 1)*(i - 1)
    }
  }
  return -1
}
