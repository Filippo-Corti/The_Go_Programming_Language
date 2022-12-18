package main

/* --- Day 5: Supply Stacks --- */

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
    [D]    
[N] [C]    
[Z] [M] [P]

This Supply Stack is represented as:
N, Z
D, C, M
P
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	supplyStack := readSupplyStack(scanner)
	supplyStack = followInstructions(scanner, supplyStack)
	fmt.Println(getItemsOnTop(supplyStack))
}

//Returns a string created by concatenating the item on top (the first item) of every row of the supplyStack
func getItemsOnTop(supplyStack [][]string) (s string) {
	for _, stack := range supplyStack {
		s += stack[0]
	}
	return
}

//Reads Input from the File until the supplyStack is finished and returns a [][]string corresponding to the supplyStack
//visualized "horizontally"
func readSupplyStack(scanner *bufio.Scanner) (supplyStack [][]string) {
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.ContainsRune(line, '['){ //End of the Supply Stack input
			return
		}
		for i, j := 1, 0; i < len(line); i, j = i + 4, j + 1 {
			//If the supplyStack doesn't have a Row for this Stack
			if len(supplyStack) <= j {
				//Create New Row of supplyStack
				supplyStack = append(supplyStack, []string{}) 
			}
			//Add item to the tail of the correct Row of the supplyStack
			//supplyStack[j] = append([]string{string(line[i])}, supplyStack[j]...) 
			if line[i] != ' ' {
				supplyStack[j] = append(supplyStack[j], string(line[i])) 
			}
		}
	}
	return
}

//Follows the Instructions read with the scanner to move the items inside the supplyStack
//And returns the new stack
func followInstructions(scanner *bufio.Scanner, supplyStack [][]string) [][]string {
	var numberOfItems, stackToMoveFrom, stackToMoveTo int
	fmt.Println("following instructions")
	//Skip first line (empty row)
	scanner.Scan()
	//Follow the instructions to move the items
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(line, "move %d from %d to %d", &numberOfItems, &stackToMoveFrom, &stackToMoveTo)
		supplyStack = moveItemsInsideStack(supplyStack, numberOfItems, stackToMoveFrom - 1, stackToMoveTo - 1)
	}
	return supplyStack
}

//Moves numberOfTimes items from the stackToMoveFrom-th row to the stackToMoveTo-th row of the supplyStack
//And returns the new stack (the indexes of the stack start from 0)
func moveItemsInsideStack(supplyStack [][]string, numberOfItems int, stackToMoveFrom int, stackToMoveTo int) [][]string {
	for i := 0; i < numberOfItems; i++ {
		itemToMove := supplyStack[stackToMoveFrom][0]
		//Add Item to The New Stack (as first item)
		supplyStack[stackToMoveTo] = append([]string{itemToMove}, supplyStack[stackToMoveTo]...)
		//Remove Item from the Old Stack
		supplyStack[stackToMoveFrom] = supplyStack[stackToMoveFrom][1:]
	}
	return supplyStack
}