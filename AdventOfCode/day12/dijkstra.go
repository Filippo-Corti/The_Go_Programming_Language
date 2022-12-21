package main

import (
	"bufio"
	"fmt"
	"os"
)

type Square struct {
	row, col int
	end bool //Is it the end?
	value rune
}



/*
I'm trying to use Dijkstra's Shortest Path algorithm on the matrix of characters of AoC Day 12
To find the shortest path from S to E.
The cost of each link is equal to 1. 
Every node is adjacent to the ones next to it that he can access
*/

func main() {
	grid, startPoint, endPoint := getGrid()
	fmt.Println(dijkstra(grid, startPoint, endPoint))
}

func dijkstra(grid [][]Square, start Square, end Square) int{
	var visited, unvisited []Square
	var dijkstraTable map[Square]int
	cols := len(grid[0])
	rows := len(grid)
	//Make all Squares unvisited initially
	visited = make([]Square, 0)
	unvisited = listItems(grid)
	//Mark all Squares with infinite distance, except for the starting square (Source)
	dijkstraTable = mapItems(grid, start)
	//Loop through every Square and update the dijkstra Table
	for len(unvisited) > 0 {
		//Get the Square with the lowest value
		lowestValued := getLowestValued(dijkstraTable, visited)
		//Mark it as visited
		visited = append(visited, lowestValued)
		//Update all adjacent vertices if the cost to reach lowestValued + 1 is less than the current cost to reach the Square to its right/left/top/bottom
		// - Right
		if lowestValued.col + 1 < cols && canIMoveTo(grid[lowestValued.row][lowestValued.col], grid[lowestValued.row][lowestValued.col + 1]) && dijkstraTable[grid[lowestValued.row][lowestValued.col]] + 1 < dijkstraTable[grid[lowestValued.row][lowestValued.col + 1]] {
			dijkstraTable[grid[lowestValued.row][lowestValued.col + 1]] = dijkstraTable[grid[lowestValued.row][lowestValued.col]] + 1
		}
		// - Left
		if lowestValued.col - 1 >= 0 && canIMoveTo(grid[lowestValued.row][lowestValued.col], grid[lowestValued.row][lowestValued.col - 1])  && dijkstraTable[grid[lowestValued.row][lowestValued.col]] + 1 < dijkstraTable[grid[lowestValued.row][lowestValued.col - 1]] {
			dijkstraTable[grid[lowestValued.row][lowestValued.col - 1]] = dijkstraTable[grid[lowestValued.row][lowestValued.col]] + 1
		}
		// - Top
		if lowestValued.row - 1 >= 0 && canIMoveTo(grid[lowestValued.row][lowestValued.col], grid[lowestValued.row - 1][lowestValued.col])  && dijkstraTable[grid[lowestValued.row][lowestValued.col]] + 1 < dijkstraTable[grid[lowestValued.row - 1][lowestValued.col]] {
			dijkstraTable[grid[lowestValued.row - 1][lowestValued.col]] = dijkstraTable[grid[lowestValued.row][lowestValued.col]] + 1
		}
		// - Bottom
		if lowestValued.row + 1 < rows && canIMoveTo(grid[lowestValued.row][lowestValued.col], grid[lowestValued.row + 1][lowestValued.col])  && dijkstraTable[grid[lowestValued.row][lowestValued.col]] + 1 < dijkstraTable[grid[lowestValued.row + 1][lowestValued.col]] {
			dijkstraTable[grid[lowestValued.row + 1][lowestValued.col]] = dijkstraTable[grid[lowestValued.row][lowestValued.col]] + 1
		}
		//Remove it from the unvisited list
		unvisited = remove(unvisited, lowestValued) 
	}
	return dijkstraTable[grid[end.row][end.col]]
}

//Checks if two squares are adjacent
func canIMoveTo(moveFrom, moveTo Square) bool {
	return int(moveTo.value) <= int(moveFrom.value) + 1
}

func remove(list []Square, s Square) []Square {
	for i, item := range list {
		if item == s {
			return append(list[:i], list[i + 1:]...)
		}
	}
	return []Square{}
}

func contains(list []Square, s Square) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}

func getLowestValued(m map[Square]int, visited []Square) Square {
	var minValue int = 99999
	var minSquare Square
	for square, value := range m {
		if value < minValue && !contains(visited, square){
			minValue = value
			minSquare = square
		}
	}
	return minSquare
} 

func mapItems(grid [][]Square, start Square) (dijkstraTable map[Square]int) {
	dijkstraTable = make(map[Square]int)
	for _, gridRow := range grid {
		for _, square := range gridRow {
			if square == start {
				dijkstraTable[square] = 0
			} else {
				dijkstraTable[square] = 999999 //Infinity
			}
		}
	}
	return
}


func listItems(grid [][]Square) (list []Square) {
	for _, gridRow := range grid {
		for _, square := range gridRow {
			list = append(list, square)
		}
	}
	return
}

//Reads the grid from the input
func getGrid() (grid [][]Square, start Square, end Square) {
	var row, col int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		gridRow := make([]Square, 0)
		line := scanner.Text()
		for _, square := range line {
			switch square {
			case 'E':
				end = Square{row, col, true, 'z'}
				gridRow = append(gridRow, end)
			case 'S':
				start = Square{row, col, false, 'a'}
				gridRow = append(gridRow, start)
			default:
				gridRow = append(gridRow, Square{row, col, false, square})
			}
			col++
		}
		grid = append(grid, gridRow)
		row++
		col = 0
	}
	return
}

func printMap(m map[Square]int) {
	for square, value := range m {
			fmt.Printf("%v: %d\n", square, value)
	}
}