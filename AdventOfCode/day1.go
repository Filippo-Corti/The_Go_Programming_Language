package main

/*
In case the Elves get hungry and need extra snacks, they need to know which Elf to ask: 
they'd like to know how many Calories are being carried by the Elf carrying the most Calories. 
Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

For example, suppose the Elves finish writing their items' Calories and end up with the following list:

1000
2000
3000

4000

5000
6000

7000
8000
9000
*/

import (
	"fmt"
	"io"
	"strconv"
)

func main() {
	var input string
	var err error
	var maxCalories, elfCalories, elfWithMaxCal, elfIndex, valuesRead int
	fmt.Println("Type CTRL-D to stop inserting data")
	for {
		for {
			valuesRead, err = fmt.Scanf("%s", &input)
			if valuesRead == 0 || err == io.EOF {
				break
			}
			calorie, _ := strconv.Atoi(input)
			elfCalories += calorie
		}
		elfIndex++
		if elfCalories > maxCalories {
			maxCalories = elfCalories
			elfWithMaxCal = elfIndex
		}
		elfCalories = 0
		if err == io.EOF {
			break
		}
	}
	fmt.Println("Max Calories:", maxCalories)
	fmt.Println("Carried by Elf number", elfWithMaxCal)
}