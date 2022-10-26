/*
Leggere un giorno, ora di inizio e fine di 2 appuntamenti e valutare se si sovrappongono
*/
package main

import "fmt"

func main() {
  var g1, h1s, h1e, g2, h2s, h2e int
  fmt.Print("Appuntamento 1 (gg, start, end): ")
  fmt.Scan(&g1, &h1s, &h1e)
  fmt.Print("Appuntamento 2 (gg, start, end): ")
  fmt.Scan(&g2, &h2s, &h2e)
  if areAppuntamentiSovrapposti(g1, h1s, h1e, g2, h2s, h2e) {
    fmt.Println("Si sovrappongono")
  } else {
    fmt.Println("Non si sovrappongono")
  }
}

func areAppuntamentiSovrapposti(g1, h1s, h1e, g2, h2s, h2e int) bool {
  if g1 != g2 {
    return false
  }
  if h2s >= h1e {
    return false
  }
  if h2e <= h1s {
    return false
  }
  return true
}
