package main

import (
  "fmt"
  "time"
)

func main() {
  fmt.Println("Divisori primi di Filippo:")
  fmt.Println(contaDivisoriPrimi(etaInGiorni(7,11,2003)))
  fmt.Println("Divisori primi di Carlotta:")
  fmt.Println(contaDivisoriPrimi(etaInGiorni(19,12,2003)))
}

/* FUNZIONI LOGICA PROGRAMMA */
  func contaDivisoriPrimi(n int) int {
    var i int
    var c int
    for i = 2; i < n; i++ {
      if n % i == 0 { //se vero i è divisore
        if isPrime(i) {
          fmt.Println(i, "è divisore primo")
          c++
        }
      }
    }
    return c
  }

  func etaInGiorni(gNascita, mNascita, aNascita int) int {
    today := time.Now()
    return distanzaFraDueDate(gNascita, mNascita, aNascita, today.Day(), int(today.Month()), today.Year())
  }

  /* FUNZIONI UTILITY */

  func isPrime (n int) bool {
    var i int
    for i = 2; i < n; i++ {
      if n % i == 0 {
        return false
      }
    }
    return true
  }

  func isAnnoBisestile (anno int) bool {
    return (anno % 100 == 0 && anno % 400 == 0) || (anno % 100 != 0 && anno % 4 == 0)
  }

  func giorniInMese (mese, anno int) int {
    if mese == 2 && isAnnoBisestile(anno) {
      return 29
    }
    if mese == 2 {
      return 28
    }
    if mese == 4 || mese == 6 || mese == 9 || mese == 11 {
      return 30
    }
    return 31
  }

  func daysFromEpoch (g, m, a int) int {
    var totale int
    //Conto Anni
    for i := 1970; i < a; i++ {
      if isAnnoBisestile(i) {
        totale += 366
        } else {
          totale += 365
        }
      }
      //Conto Mesi
      for i := 1; i < m; i++ {
        totale += giorniInMese(i,a)
      }
      //Conto Giorni
      totale += g
      //Restituisco il totale
      return totale
    }

  func distanzaFraDueDate(g1, m1, a1, g2, m2, a2 int) int {
      giorni1 := daysFromEpoch(g1,m1,a1)
      giorni2 := daysFromEpoch(g2,m2,a2)
      return giorni2 - giorni1
  }
