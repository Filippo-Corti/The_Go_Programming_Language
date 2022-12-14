/*
Scrivere un programma che riceve in ingresso due numeri float64 rappresentanti
peso e altezza e stampa un responso in funziond della tabella
*/

package main

import "fmt"

func main() {
  var peso, altezza float64
  fmt.Print("Peso (kg) Altezza (cm):")
  fmt.Scan(&peso, &altezza)
  valutaStatoSalute(peso, altezza)
}

func valutaStatoSalute(peso, altezza float64) {
  if peso >= 10 && peso <= 50 {
    if altezza >= 100 && altezza <= 150 {
      fmt.Println("normopeso")
      return
    }
    if altezza > 150 && altezza <= 200 {
      fmt.Println("sottopeso")
      return
    }
  }
  if peso > 50 && peso <= 100 {
    if altezza >= 100 && altezza <= 150 {
      fmt.Println("sovrappeso")
      return
    }
    if altezza > 150 && altezza <= 200 {
      fmt.Println("normopeso")
      return
    }
  }
  if peso > 100 {
    if altezza >= 100 && altezza <= 150 {
      fmt.Println("molto sovrappeso")
      return
    }
    if altezza > 150 && altezza <= 200 {
      fmt.Println("sovrappeso")
      return
    }
  }

}
