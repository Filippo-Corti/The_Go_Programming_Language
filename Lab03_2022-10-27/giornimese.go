package main

import "fmt"

func main() {
  var mese int
  fmt.Print("Inserire intero da 1 a 12: ")
  fmt.Scan(&mese)
  fmt.Println("Il mese ha", giorniInMese(mese), "giorni")
}

func giorniInMese(mese int) int {
  if mese == 2 {
    return 28
  }
  if mese == 11 || mese == 4 || mese == 6 || mese == 9 {
    return 30
  }
  return 31
}
