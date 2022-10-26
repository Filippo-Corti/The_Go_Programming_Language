/*
Controllo data valida (giorno mese)
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
    if day == 31 && !isMonth31DaysLong(month)  {
      return false
    }
    if day > 28 && month == 2 {
      return false
    }
    return true
}

func isMonth31DaysLong(month int) bool {
  return month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12
}
