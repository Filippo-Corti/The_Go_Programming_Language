package main
import (
    "fmt"
    "math/rand"
    "time"
)
func main() {
    /*
		Genera numeri casuali fra 0 e K fino a che non genera 0, poi stampa i tentativi 
		impiegati 
	*/
    rand.Seed(time.Now().Unix())
    const K = 20
    var n int

    c := 0
    for {
        n = rand.Intn(K+1)
        fmt.Print(n, " ")
        if n == 0 {
            break
        }
        c++
    }
    fmt.Println(c)
}
