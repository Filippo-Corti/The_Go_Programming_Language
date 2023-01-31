/*Codice del programma estraiData*/
package main

import (
   "fmt"
   "strings"
   "strconv"
)

func main() {
   var dataGMA string
   fmt.Print("data nel formato giorno/mese/anno: ")
   fmt.Scan(&dataGMA)
   g, m, a := stringGMA2GMA(dataGMA)
   fmt.Println("giorno", g)
   fmt.Println("mese", num2mese(m))
   fmt.Println("anno", a)

}

func stringGMA2GMA(data string) (int, int, int) {
	firstSlash := strings.Index(data, "/")
	lastSlash := strings.LastIndex(data, "/")
	giorno, _ := strconv.Atoi(data[: firstSlash])
	mese, _ := strconv.Atoi(data[firstSlash + 1 : lastSlash])
	anno, _ := strconv.Atoi(data[lastSlash + 1 :])
   return giorno, mese, anno
}

func num2mese(m int) string {
   var mesi = [13]string{"", "gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"}
   return mesi[m]
}