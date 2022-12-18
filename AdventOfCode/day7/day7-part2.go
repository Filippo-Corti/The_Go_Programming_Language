package main

/* --- Day 7: No Space Left On Device --- */

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	name string
	size int //0 if it's a directory
	next []*Node
	prev *Node
}

func main() {
	var correctDirSize int = 30000000
	instructions := getInstructions()
	fileSystem := getFileSystem(instructions)
	totalUsedSpace := calculateDirectoryDimensions(fileSystem)
	totalUnusedSpace := 70000000 - totalUsedSpace
	extraSpaceNeeded := 30000000 - totalUnusedSpace
	findCorrectDir(fileSystem, extraSpaceNeeded, &correctDirSize)
	fmt.Println(correctDirSize)
}

func findCorrectDir(fileSystem *Node, extraSpaceNeeded int, correctDirSize *int) (sum int) {
	for _, node := range fileSystem.next {
		if node.size == 0 {
			//It's a Directory
			relativeDim := findCorrectDir(node, extraSpaceNeeded, correctDirSize)
			sum += relativeDim
		} else {
			sum += node.size
		}
	}
	if sum < *correctDirSize && sum >= extraSpaceNeeded {
		*correctDirSize = sum
	}
	return sum
}

func calculateDirectoryDimensions(fileSystem *Node) (dim int) {
	for _, node := range fileSystem.next {
		if node.size == 0 {
			//It's a Directory 
			partialDim := calculateDirectoryDimensions(node) 
			dim += partialDim
		} else {
			//It's a File
			dim += node.size
		}
	}
	return
}

func getFileSystem(instr []string) *Node {
	var current *Node
	i := 1 //Skip "$ cd /"
	slash := new(Node)
	slash.name = "/"
	current = slash
	for i < len(instr) {
		parameters := strings.Split(instr[i], " ")
		if parameters[0] == "$" {
			//Instruction is a command
			switch parameters[1] {
			case "cd":
				//Move inside one node
				//fmt.Println("cd")
				current = changeDir(current, parameters[2])
			case "ls":
				//Create new subnodes inside current.next
				//fmt.Println("ls")
				instr = listDirectories(instr[i + 1:], current)
				i = -1
			}
		}
		i++
	}
	return slash
}

func changeDir(current *Node, destination string) *Node {
	if destination == ".." {
		return current.prev
	}
	for _, node := range current.next {
		if node.name == destination {
			return node
		}
	}
	return current
}


func listDirectories(instr []string, current *Node) []string {
	var i int
	for i = 0; i < len(instr) && instr[i][0] != '$' ; i++ {
		//Create New Node
		newNode := new(Node)
		parameters := strings.Split(instr[i], " ")
		if parameters[0] != "dir" {
			//It's a file
			newNode.size, _ = strconv.Atoi(parameters[0])
		}
		//If it's a directory, size remains 0
		newNode.name = parameters[1]
		newNode.prev = current
		current.next = append(current.next, newNode)
	}
	return instr[i:]
}

func getInstructions() (instructions []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, line)
	}
	return
}