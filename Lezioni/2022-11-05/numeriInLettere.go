package main

import "fmt"

var DIECI int = 10
var CENTO int = 100
var MILLE int = 1000
var MILIONE int = 1000000
var MILIARDO int = 1000000000

func main() {
	var n int
	fmt.Print("Number: ")
	fmt.Scan(&n)
	fmt.Println("In Letters:", numeroInLettere(n))
}

func numeroInLettere(n int) string {
	if n >= 1 && n <= 9 {
		return numeroInLettere1To9(n)
	}
	if n >= DIECI && n <= (CENTO - 1) {
		return numeroInLettere10To99(n)
	}
	if n >= CENTO && n <= (MILLE - 1) {
		return numeroInLettere100To999(n)
	}
	if n >= MILLE && n <= (MILIONE - 1) {
		return numeroInLettere1000To999999(n)
	}
	if n >= MILIONE && n <= (MILIARDO - 1) {
		return numeroInLettere1000000To999999999(n) 
	}
	if n >= MILIARDO {
		return numeroInLettere1000000000To999999999999(n)
	}
	return ""
}

func numeroInLettere1To9(n int) string {
	switch n {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	default:
		return ""
	}
}

func numeroInLettere10To99(n int) string {
	ris := ""
	switch n / 10 {
	case 1:
		return numeroInLettere10To19(n)
	case 2:
		ris += "twenty"
	case 3:
		ris += "thirty"
	case 4:
		ris += "forty"
	case 5:
		ris += "fifty"
	case 6:
		ris += "sixty"
	case 7:
		ris += "seventy"
	case 8:
		ris += "eighty"
	case 9:
		ris += "ninety"
	}
	if n%10 == 0 {
		return ris
	} else {
		ris += "-" + numeroInLettere1To9(n%10)
	}
	return ris
}

func numeroInLettere10To19(n int) string {
	switch n {
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "fifteen"
	case 16:
		return "sixteen"
	case 17:
		return "seventeen"
	case 18:
		return "eighteen"
	case 19:
		return "nineteen"
	default:
		return ""
	}
}

func numeroInLettere100To999(n int) string {
	ris := numeroInLettere1To9(n / 100) + " hundred " 
	if n % 100 < 10 {
		ris += numeroInLettere1To9(n % 100)
	} else {
		ris += numeroInLettere10To99(n % 100)
	}
	return ris
}

func numeroInLettere1000To999999(n int) string {
	return numeroInLettere(n / MILLE) + " thousand " + numeroInLettere(n % MILLE)
}

func numeroInLettere1000000To999999999(n int) string {
	return numeroInLettere(n / MILIONE) + " million " + numeroInLettere(n % MILIONE)
}

func numeroInLettere1000000000To999999999999(n int) string {
	return numeroInLettere(n / MILIARDO) + " billion " + numeroInLettere(n % MILIARDO)
}
