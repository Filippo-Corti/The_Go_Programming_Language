package main

/*

*/

import (
	"fmt"
	"io"
	"sort"
	"strconv"
)

func main() {
	var input string
	var err error
	var elfCalories, elfIndex, valuesRead int
	var top3Calories []int
	top3Calories = make([]int, 0)
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
		//Fill in the list with the first 3 elves
		if len(top3Calories) < 3 {
			top3Calories = append(top3Calories, elfCalories)
			//Sort the list
			sort.Ints(top3Calories)
		}
		if elfCalories > top3Calories[0] {
			//Add to the list
			top3Calories = append(top3Calories, elfCalories)
			//Sort the list
			sort.Ints(top3Calories)
			//Remove the smalles number
			top3Calories = top3Calories[1:]
		}
		elfCalories = 0
		if err == io.EOF {
			break
		}
	}
	fmt.Println("Max Calories:", sum(top3Calories))
}

func sum(list []int) (s int){
	for _, item := range list {
		s += item
	}
	return
}