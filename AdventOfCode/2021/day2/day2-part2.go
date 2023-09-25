package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	horizontalPosition, finalDepth := getData()
	fmt.Println(horizontalPosition * finalDepth)
}

func getData() (horizontalPosition, finalDepth int) {
	var aim int
	var command string
	var value int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(line, "%s %d\n", &command, &value)
		switch(command) {
		case "forward":
			horizontalPosition += value
			finalDepth += value * aim
		case "up": 
			aim -= value
		case "down":
			aim += value
		}
	} 
	return 
}