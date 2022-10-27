package main

import (
  "fmt"
  "math"
)

func main() {
  var bin int
  fmt.Print("Inserire numero binario: ")
  fmt.Scan(&bin)
  fmt.Println("Corrispettivo in decimale:", binToDec(bin))
}

func binToDec(b int) (d int) {
  c := 0
  for {
    cifra := b % 10
    b /= 10
    if cifra != 0 && cifra != 1 {
      return -1
    }
    d += cifra * int(math.Pow(2.0,float64(c)))
    c++
    if b == 0 {
      return d
      break
    }
  }
  return -1
}
