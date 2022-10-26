/* Leggere i due vertici opposti di un rettangolo e valutare se un terzo punto è al suo interno */
package main

import "fmt"

func main() {
  var x1, y1, x2, y2, x3, y3 int
  fmt.Print("P1: ")
  fmt.Scan(&x1, &y1)
  fmt.Print("P2: ")
  fmt.Scan(&x2, &y2)
  fmt.Print("P3: ")
  fmt.Scan(&x3, &y3)

  if valoreCompresoTraAltriDue(x3, x1, x2, false) && valoreCompresoTraAltriDue(y3, y1, y2, false)  {
    fmt.Println("P3 Interno")
  } else if valoreCompresoTraAltriDue(x3, x1, x2, true) && valoreCompresoTraAltriDue(y3, y1, y2, true){
    fmt.Println("P3 Perimetrale")
  }else {
    fmt.Println("P3 Esterno")
  }
}

// IntervalloChiuso : TRUE = [b,c] - FALSE = (b,c)
func valoreCompresoTraAltriDue(a, b, c int, intervalloChiuso bool) bool {
  if c < b {
    b, c = c, b
  }
  if intervalloChiuso {
      return (a >= b && a <= c)
  } else {
    return (a > b && a < c)
  }
}
