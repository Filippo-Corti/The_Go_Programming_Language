package main

import "fmt"

var HEARTS rune = '\u2661'   //Cuori ♡
var DIAMONDS rune = '\u2662' //Quadri ♢
var CLUBS rune = '\u2667'    //Fiori ♧
var SPADES rune = '\u2664'   //Picche ♤

func main() {
	printActualDeck()
}

func printActualDeck() {
	for j := 0; j < 4; j++ {
		for i := 0; i < 13; i++ {
			fmt.Print(string('\U0001F0A1' + j * 16 + i) + " ")
		}
		fmt.Println()
	}
}

func printDeck() {
	for j := 0; j < 4; j++ {
		for i := 0; i < 13; i++ {
			printCard(j*13 + i)
		}
		fmt.Println()
	}
}

func printCard(n int) {
	if n < 0 || n > 51 {
		fmt.Println("\n!ERROR - Card Out of Range!\n")
		return
	}
	fmt.Print(getCardSuit(n), " ",  getCardValue(n), "\t")
}

func getCardValue(n int) string {
	pos := n%13 + 1
	switch {
	case pos > 1 && pos < 10:
		return string(pos + 48)
	case pos == 1:
		return "A"
	case pos == 10:
		return "10"
	case pos == 11:
		return "J"
	case pos == 12:
		return "Q"
	case pos == 13:
		return "K"
	}
	return ""
}

func getCardSuit(n int) string {
	switch n / 13 {
	case 0:
		return string(HEARTS)
	case 1:
		return string(DIAMONDS)
	case 2:
		return string(CLUBS)
	case 3:
		return string(SPADES)
	}
	return ""
}
