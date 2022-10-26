/*
Controllo data valida con mesi di tutti 31 giorni
*/
package main

import "fmt"

func main() {
  var g, m int
  fmt.Print("gg mm? ")
  fmt.Scan(&g, &m)
  if isDataValida(g, m){
    fmt.Println("Data Valida")
  } else {
    fmt.Println("Data non Valida")
  }
}

func isDataValida(day, month int) bool {
    if day < 1 || day > 31 {
      return false
    }
    if month < 1 || month > 12 {
      return false
    }
    return true
}
