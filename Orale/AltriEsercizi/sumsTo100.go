// Print all sums and subtractions of numbers 1 to 9 (in this order) to get 100.
// es: 1 + 2 + 3 - 4 + 5 + 6 + 78 + 9 = 100
package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type Possibility struct {
	expression string // 1 + 2 + 3 + 4 - 5 + 67 ...
}

func (p Possibility) Calculate() (sum int) {
	r1, _ := regexp.Compile("[0-9]+")
	numbers := r1.FindAllString(p.expression, -1)
	r2, _ := regexp.Compile("[+-]")
	operators := r2.FindAllString(p.expression, -1)
	sum, _ = strconv.Atoi(numbers[0])
	for i := 0; i < len(operators); i++ {
		number, _ := strconv.Atoi(numbers[i + 1])
		if operators[i] == "+" {
			sum += number
		} else {
			sum -= number
		}
	}
	return
}

func main() {
	fmt.Println("Esiste sicuramente un modo più efficiente ma questo mi piaceva molto :)")
	possibilities := generatePossibilities(9)
	for _, possibility := range possibilities {
		if possibility.Calculate() == 100 {
			fmt.Println(possibility)
		}
	}
}

func generatePossibilities(newNumber int) (possibilities []Possibility) {
	//Caso base
	if newNumber == 1 {
		possibilities = append(possibilities, Possibility{"1"})
		return
	}
	previousPossibilities := generatePossibilities(newNumber - 1)
	for _, previousPossibility := range previousPossibilities {
		possibilities = append(possibilities, newPossibilityBySum(previousPossibility, newNumber))
		possibilities = append(possibilities, newPossibilityBySubtraction(previousPossibility, newNumber))
		possibilities = append(possibilities, newPossibilityByCoupling(previousPossibility, newNumber))
	}
	return
}

func newPossibilityBySum(old Possibility, n int) (new Possibility) {
	new.expression = old.expression + fmt.Sprintf(" + %d", n)
	return
}

func newPossibilityBySubtraction(old Possibility, n int) (new Possibility) {
	new.expression = old.expression + fmt.Sprintf(" - %d", n)
	return
}

func newPossibilityByCoupling(old Possibility, n int) (new Possibility) {
	new.expression = old.expression + fmt.Sprintf("%d", n)
	return
}