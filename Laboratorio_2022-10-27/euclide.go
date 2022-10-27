package main

import "fmt"

func main() {
  var a, b int
  fmt.Print("Inserire a e b: ")
  fmt.Scan(&a, &b)
  fmt.Println("MCD:", mcdConEuclide(a, b))
}

func mcdConEuclide(a, b int) (mcd int) {
  for b != 0 {
    a, b = b, a % b
  }
  return a
}
