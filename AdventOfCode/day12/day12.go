package main

/* --- Day 12: Hill Climbing Algorithm --- */

import (
	"bufio"
	"fmt"
	"os"
)

type Square struct {
	row, col int
	end bool //Is it the end?
	value rune
	visited bool
}

type Path struct {
	direction string
	steps int
}

func main() {
	grid, startPoint := getGrid()
	bestPath, _ := countShortestPath(grid, startPoint)
	//printGrid(grid)
	fmt.Println(bestPath)
}

//Receives a grid and a starting point and returns the minimum number of steps to get to the end and the path it takes to get there
func countShortestPath(grid [][]Square, start Square) (steps int, visited []Square) {
	//fmt.Println(start)
	var visitedRight, visitedLeft, visitedUp, visitedDown []Square
	var right, left, up, down int
	visited = make([]Square, 0)
	pathCounts := make([]Path, 0)
	rows := len(grid)
	cols := len(grid[0])
	//I track that I've been here
	visited = append(visited, start)
	grid[start.row][start.col].visited = true
	//If I'm standing on the 'E' path has lenght of zero
	if grid[start.row][start.col].end == true {
		return 0, visited
	}
	//In this case instead, I need to find the fastest way 
	//Try Right:
	if start.col + 1 < cols && shouldIMoveTo(start, grid[start.row][start.col + 1]) {
		right, visitedRight = countShortestPath(grid, grid[start.row][start.col + 1])
		if right != -1 {
			//Going Right I get to the end
			pathCounts = append(pathCounts, Path{"R", right})
		}
		clearPath(grid, visitedRight)
	}
	//Try Left:
	if start.col - 1 >= 0 && shouldIMoveTo(start, grid[start.row][start.col - 1]) {
		left, visitedLeft = countShortestPath(grid, grid[start.row][start.col - 1]) 
		if left != -1 {
			//Going Left I get to the end
			pathCounts = append(pathCounts, Path{"L", left})
		}
		clearPath(grid, visitedLeft)
	}
	//Try Up:
	if start.row - 1 >= 0 && shouldIMoveTo(start, grid[start.row - 1][start.col]) {
		up, visitedUp = countShortestPath(grid, grid[start.row - 1][start.col]) 
		if up != -1 {
			//Going Up I get to the end
			pathCounts = append(pathCounts, Path{"U", up})
		}
		clearPath(grid, visitedUp)
	}
	//Try Down:
	if start.row + 1 < rows && shouldIMoveTo(start, grid[start.row + 1][start.col]) {
		down, visitedDown = countShortestPath(grid, grid[start.row + 1][start.col]) 
		if down != -1 {
			//Going Up I get to the end
			pathCounts = append(pathCounts, Path{"D", down})
		}
		clearPath(grid, visitedDown)
	}
	//Now clearPath contains every valuable Path
	//I first need to check that it's not empty
	if len(pathCounts) == 0 {
		//This is NOT the right way
		//fmt.Println("We should go back...", start)
		return -1, visited
	}
	bestPath := getTheBestPath(pathCounts)
	//I need to make sure to add the path of the direction I took to my "visited" path
	switch bestPath.direction {
	case "R":
		visited = append(visited, visitedRight...)
		//fmt.Println("Going Right from", start)
	case "L":
		visited = append(visited, visitedLeft...)
		//fmt.Println("Going Left from", start)
	case "U":
		visited = append(visited, visitedUp...)
		//fmt.Println("Going Up from", start)
	case "D":
		visited = append(visited, visitedDown...)
		//fmt.Println("Going Down from", start)
	}
	//I can now return the best possible path starting from "Start"
	return bestPath.steps + 1, visited
}

//It receives a list of possible Paths and returns the best one (shortest one)
func getTheBestPath(pathCounts []Path) Path {
	bestPath := pathCounts[0]
	for i := 1; i < len(pathCounts); i++ {
		if bestPath.steps > pathCounts[i].steps {
			bestPath = pathCounts[i]
		}
	}
	return bestPath
}

//Sets all the elements of visited to unvisited
func clearPath(grid [][]Square, visited []Square) {
	//I need to clear the visited list from the spots I covered moving right
	//This way I can move to any other direction with the same grid situation as before
	for _, square := range visited {
		grid[square.row][square.col].visited = false
	}
}

//Checks if I've already been there or if the jump is not possible
func shouldIMoveTo(current, moveTo Square) bool {
	return !moveTo.visited && canIMoveTo(current, moveTo)
}

//Checks if the jump is possible 
func canIMoveTo(moveFrom, moveTo Square) bool {
	return int(moveTo.value) <= int(moveFrom.value) + 1
}

//Reads the grid from the input
func getGrid() (grid [][]Square, start Square) {
	var row, col int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		gridRow := make([]Square, 0)
		line := scanner.Text()
		for _, square := range line {
			switch square {
			case 'E':
				gridRow = append(gridRow, Square{row, col, true, 'z', false})
			case 'S':
				start = Square{row, col, false, 'a', true}
				gridRow = append(gridRow, Square{row, col, false, 'a', true})
			default:
				gridRow = append(gridRow, Square{row, col, false, square, false})
			}
			col++
		}
		grid = append(grid, gridRow)
		row++
		col = 0
	}
	return
}

//Prints the grid
func printGrid(grid [][]Square) {
	for _, gridRow := range grid {
		for _, square := range gridRow {
			if square.end == true {
				fmt.Printf("E")
				continue
			}
			//fmt.Printf("%c", square.value)
			if square.visited {
				fmt.Printf("V")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
} 