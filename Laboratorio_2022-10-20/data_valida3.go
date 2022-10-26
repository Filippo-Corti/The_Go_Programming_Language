/*
Controllo data valida (giorno mese)
*/
package main

import "fmt"

func main() {
  var g, m, a int
  fmt.Print("gg mm aaaa? ")
  fmt.Scan(&g, &m, &a)
  if isDataValida(g, m, a){
    fmt.Println("Data Valida")
  } else {
    fmt.Println("Data non Valida")
  }
}

func isDataValida(day, month, year int) bool {
    if day < 1 || day > 31 {
      return false
    }
    if month < 1 || month > 12 {
      return false
    }
    if day == 31 && !isMonth31DaysLong(month)  {
      return false
    }
    if month == 2 && (day > 29 && isBisestile(year) || day > 28 && !isBisestile(year)) {
      return false
    }
    return true
}

func isMonth31DaysLong(month int) bool {
  return month == 1 || month == 3 || month == 5 || month == 7 || month == 8 || month == 10 || month == 12
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
