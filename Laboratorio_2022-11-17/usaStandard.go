package main

import "fmt"
import "strings"
import "strconv"
import "math"

var VOCALI string = "aeiouAEIOU"
var CIFRE string = "0123456789"
var S1 string = "c"
var S2 string= "sc"

func main() {
	var s string
	fmt.Print("Una stringa: ")
	fmt.Scan(&s)
	if strings.Contains(s, S2) {
		fmt.Printf("1.\t%s contiene %s\n",s, S2)
	} else {
		fmt.Printf("1.\t%s non contiene %s\n",s, S2)
	}
	if strings.ContainsAny(s, VOCALI) {
		fmt.Printf("2.\t%s contiene vocali\n",s)
	} else {
		fmt.Printf("2.\t%s non contiene vocali\n",s)
	}
	fmt.Printf("3.\t%s contiene %d %s\n", s, strings.Count(s, S1), S1)
	fmt.Printf("4.\tLa prima vocale di %s si trova in posizione %d\n", s, strings.IndexAny(s, VOCALI))
	fmt.Printf("5.\tL'ultima vocale di %s si trova in posizione %d\n", s, strings.LastIndexAny(s, VOCALI))
	fmt.Printf("6.\t%s\n", strings.Repeat(S2, 3))
	numeri, _ := strconv.Atoi(extractNumbers(s))
	fmt.Printf("7.\t Int %d\n", numeri)
	fmt.Printf("8.\t Int %f\n", float64(numeri)/math.Pow(10.0, float64(len(strconv.Itoa(numeri)))))
}

func extractNumbers(s string) (numeri string) {
	for _,r := range s {
		if strings.ContainsRune(CIFRE, r){
			numeri += string(r)
		}
	}
	return
} 