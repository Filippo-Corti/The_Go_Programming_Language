package main

import (
	"fmt"
	"math/rand"
)

const MAX = 10

func main() {
	var chosen, prec int
	var typed string
	var wins, losses int
	prec = rand.Intn(MAX) + 1
	for {
		for chosen == prec { //To Avoid the situation where they are equal and the player can only lose
			chosen = rand.Intn(MAX) + 1
		}
		fmt.Print("alto/basso? ")
		fmt.Scan(&typed)
		if typed == "0" {
			break
		}
		if typed == "alto" && chosen > prec || typed == "basso" && chosen < prec {
			wins++
			fmt.Println("Corretto")
		} else {
			losses++
			fmt.Println("Sbagliato")
		}
		prec = chosen
	}
	fmt.Printf("Vittorie: %.2f %%\n", float64(wins)/float64(wins + losses)*100)
	fmt.Printf("Sconfitte: %.2f %%\n", float64(losses)/float64(wins + losses)*100)
}