package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	depths := parseInput()
	fmt.Println(countIncrements(depths))
}


func countIncrements(depths []int) (count int) {
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i - 1] {
			count++
		}
	}
	return
}

func parseInput() (depths []int) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		depthString := scanner.Text()
		depthInt, _ := strconv.Atoi(depthString)
		depths = append(depths, depthInt)
	}
	return
}
