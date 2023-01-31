package main
import "fmt"
func main() {
    /*
		Leggere un intero n e stampare i numeri pari in [0,n]
	*/
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 0; i <= n; i = i + 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
