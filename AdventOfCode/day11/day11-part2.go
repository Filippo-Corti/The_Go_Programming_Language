package main

/* --- Day 11: Monkey in the Middle --- */

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sort"
)

type Monkey struct {
	items []int
	operation func(int)int
	test int
	nextMonkeyIfTrue int
	nextMonkeyIfFalse int
	itemsInspected int
}


func main() {
	monkeys := getMonkeys()
	// for _, monkey := range monkeys {
	// 	fmt.Println(monkey.operation(1))
	// }
	playMonkeyInTheMiddle(monkeys)
	fmt.Println(monkeys)
	fmt.Println("Monkey Business:", getMonkeyBusiness(monkeys))
}

func getMonkeyBusiness(monkeys []Monkey) int {
	top2 := []int{}
	for i := 0; i < len(monkeys); i++ {
		top2 = append(top2, monkeys[i].itemsInspected)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(top2)))
	fmt.Println(top2)
	return top2[0] * top2[1]
}

func playMonkeyInTheMiddle(monkeys []Monkey) int {
	MCDofAllTests := getMCDofTests(monkeys)
	//Play for 10000 rounds
	for i := 0; i < 10000; i++ {
		for i, _ := range monkeys {
			//Play monkey turn
			playMonkeyTurn(&monkeys[i], monkeys, MCDofAllTests)
			//fmt.Println(monkeys)
		}
	}
	return 0
}

func getMCDofTests(monkeys []Monkey) int {
	mcd := 1
	for _, monkey := range monkeys {
		mcd *= monkey.test
	}
	fmt.Println(mcd)
	return mcd
}

func playMonkeyTurn(monkey *Monkey, monkeys []Monkey, MCDofAllTests int) {
	for i, _ := range monkey.items {
		//Change the worry level of the item
		//fmt.Printf("Monkey inspects an item with a worry level of %d.\n", monkey.items[i])
		monkey.items[i] = monkey.operation(monkey.items[i]) % MCDofAllTests
		//fmt.Printf("Worry level is changed to %d.\n", monkey.items[i])
		//monkey.items[i] = divideByThreeApprox(monkey.items[i])
		//fmt.Printf("Worry level is divided by 3 to %d.\n", monkey.items[i])
		monkey.itemsInspected++
		//Pass it to the next monkey 
		if monkey.items[i] % monkey.test == 0 {
			//fmt.Printf("Worry level is divisible by %d\nItem %d is thrown to Monkey %d.\n", monkey.test, monkey.items[i], monkey.nextMonkeyIfTrue)
			monkeys[monkey.nextMonkeyIfTrue].items = append(monkeys[monkey.nextMonkeyIfTrue].items, monkey.items[i])
		} else {
			//fmt.Printf("Worry level is not divisible by %d\nItem %d is thrown to Monkey %d.\n", monkey.test, monkey.items[i], monkey.nextMonkeyIfFalse)
			monkeys[monkey.nextMonkeyIfFalse].items = append(monkeys[monkey.nextMonkeyIfFalse].items, monkey.items[i])
		}
	}
	monkey.items = []int{}
}

func divideByThreeApprox(n int) int {
	return int(math.Floor(float64(n)/3.0))
}

func getMonkeys() (monkeys []Monkey) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() { //Monkey 0:
		monkey := new(Monkey) 
		scanner.Scan()
		startingItems := scanner.Text() //Starting items: 79, 98
		monkey.items = getItems(startingItems)
		scanner.Scan()
		operation := scanner.Text() //Operation: new = old * 19
		monkey.operation = getOperation(operation)
		scanner.Scan()
		test := scanner.Text() //Test: divisible by 23
		monkey.test = getTest(test)
		scanner.Scan()
		ifTrue := scanner.Text() //If true: throw to monkey 2
		monkey.nextMonkeyIfTrue = getIfTrue(ifTrue)
		scanner.Scan()
		ifFalse := scanner.Text() //If false: throw to monkey 3
		monkey.nextMonkeyIfFalse = getIfFalse(ifFalse)
		monkeys = append(monkeys, *monkey)
		scanner.Scan() //Empty line
	}
	return
}

func getFunction(operation rune, value int, isValueOld bool) func(int)int {
	if isValueOld {
		return func(n int) int {
			return n * n
		}
	}
	if operation == '+' {
		return func(n int) int {
			return n + value
		}
	}
	if operation == '*'{
		return func(n int) int {
			return n * value
		}
	}
	return nil
}

func getIfFalse(s string) int {
	s = strings.TrimPrefix(s,"    If false: throw to monkey ")
	value, _ := strconv.Atoi(s)
	return value
}

func getIfTrue(s string) int {
	s = strings.TrimPrefix(s,"    If true: throw to monkey ")
	value, _ := strconv.Atoi(s)
	return value
}

func getTest(s string) int {
	s = strings.TrimPrefix(s,"  Test: divisible by ")
	value, _ := strconv.Atoi(s)
	return value
}

func getOperation(s string) func(int)int {
	s = strings.TrimPrefix(s, "  Operation: new = old ")
	operationFields := strings.Fields(s)
	value, err := strconv.Atoi(operationFields[1])
	if err != nil {
		return getFunction(rune(operationFields[0][0]), -1, true)
	} else {
		return getFunction(rune(operationFields[0][0]), value, false)
	}
}

func getItems(s string) (itemsInt []int) {
	s = strings.TrimPrefix(s, "  Starting items: ")
	items := strings.Split(s, ",")
	for _, item := range items {
		item = strings.TrimSpace(item)
		value, _ := strconv.Atoi(item)
		itemsInt = append(itemsInt, value)
	}
	return
}	