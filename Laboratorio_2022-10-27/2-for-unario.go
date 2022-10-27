package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
    /*
		Genera e stampa numeri random da 1 a MAX fino a che la
		somma dei valori generati supera TARGET. Dopodiché 
		stampa il totale raggiunto
    */
    const TARGET = 20
    const MAX = 10
    rand.Seed(time.Now().Unix())
    var n int

    totale := 0
    for totale < TARGET {
        n = rand.Intn(MAX) + 1
        fmt.Print(n, " ")
        totale += n
    }
    fmt.Println(totale)
}
