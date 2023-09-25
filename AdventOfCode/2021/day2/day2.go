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
	var command string
	var value int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(line, "%s %d\n", &command, &value)
		switch(command) {
		case "forward":
			horizontalPosition += value
		case "up": 
			finalDepth -= value
		case "down":
			finalDepth += value
		}
	} 
	return 
}