/* Leggere coordinate e valutare in che quadrante si trova il punto */
package main

import "fmt"

func main() {
  var x, y float64
  fmt.Print("x y: ")
  fmt.Scan(&x, &y)
  if x >= 0 {
    if y >= 0 {
      fmt.Println("I Quadrante")
    } else {
      fmt.Println("IV Quadrante")
    }
  } else {
    if y >= 0 {
      fmt.Println("II Quadrante")
    } else {
      fmt.Println("III Quadrante")
    }
  }
}
