package main
import "fmt"
func main() {
    /*
	 Scrivere un programma che legga un numero n e stampi n asterischi, per poi andare a capo
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}
