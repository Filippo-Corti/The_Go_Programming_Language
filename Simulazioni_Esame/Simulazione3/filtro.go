package main

import (
	"fmt"
)

const INTERVALS = 4

func main() {
	var counters [INTERVALS]int
	var num, total int
	for {
		fmt.Scanf("%d", &num)
		//Stop at '0'
		if num == 0 {
			break
		}
		//Ignore numbers outside [18,30]
		if num < 18 || num > 30 {
			continue
		}
		counters[rangeOf(num)]++
		total++
	}
	printResults(counters, total)
}

func printResults(counters [INTERVALS]int, total int) {
	fmt.Println(counters, total)
	for i, c := INTERVALS - 1, 'A'; i >= 0; i, c = i - 1, c + 1 {
		fmt.Printf("%c : %d %%\n", c, int((float64(counters[i]) / float64(total)) * 100))
	}
}

//Returns the index of the range corresponding to n
//Ranges are: [D: 18-21, C: 22-24, B: 25-27, A:28-30]
func rangeOf(n int) int {
	return (n - 19) / 3
}