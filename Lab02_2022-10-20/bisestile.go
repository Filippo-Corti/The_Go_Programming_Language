/* Valutare se anno è bisestile */
package main

import "fmt"

func main() {
  var a int
  fmt.Print("anno? ")
  fmt.Scan(&a)
  if isBisestile(a){
    fmt.Println("Bisestile")
  } else {
    fmt.Println("Non bisestile")
  }
}

func isBisestile(year int) bool {
  if year < 4 {
    return false
  }
  if year >= 4 && year <= 1580 {
    return year % 4 == 0
  }
  return (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0) 
}
