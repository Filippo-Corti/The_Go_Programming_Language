package main
import "fmt"
func main() {
    /*
		Leggere un numero n in input e stampare i numeri da n a 1 (compresi)
	*/
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := n; i >= 1; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}
