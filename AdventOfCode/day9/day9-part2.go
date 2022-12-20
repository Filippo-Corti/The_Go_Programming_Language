package main

/* --- Day 9: Rope Bridge --- */

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Knot struct {
	name string
	row, col int
}

type Direction struct {
	orientation string
	distance int
}

func main() {
	directions := getDirections()
	bridge := getBridgeGrid(directions)
	followDirections(bridge, directions)
	fmt.Println(countVisitedCells(bridge))
}

func countVisitedCells(bridge [][]string) (c int) {
	for _, bridgeRow := range bridge {
		for _, value := range bridgeRow {
			if value == "#" {
				c++
			}
		}
	}
	return
}

func followDirections(bridge [][]string, directions []Direction) () {
	_, _, startRow, startCol := getGridDimension(directions)
	tails := make([]Knot, 9) 
	for i := 0; i < len(tails); i++ {
		tails[i] = Knot{strconv.Itoa(i + 1), startRow, startCol}
	}
	head := Knot{"H", startRow, startCol}
	for _, direction := range directions {
		for i := 0; i < direction.distance; i++ {
			moveHead(&head, direction.orientation)
			moveTail(&tails[0], head)
			for i := 1; i < len(tails); i++ {
				moveTail(&tails[i], tails[i - 1])
			}
			//fmt.Println(head.row, head.col, tails)
			bridge[tails[8].row][tails[8].col] = "#"
			//printGrid(bridge, tails, head)
		}
	}
}

func moveTail(tail *Knot, head Knot) {
	if tail.row == head.row || tail.col == head.col {
		//They are on the same row or col
		if tail.row == head.row && tail.col == head.col {
			//They are on the same spot: no moving
			return
		} 
		if tail.row == head.row {
			//They are on the same row:
			moveKnot(tail, tail.row, tail.col + (head.col - tail.col) / 2)
		} else {
			//They are on the same col
			moveKnot(tail, tail.row + (head.row - tail.row) / 2, tail.col)
		}
	} else if math.Abs(float64(tail.row - head.row)) > 1.0 || math.Abs(float64(tail.col - head.col)) > 1.0 {
		if math.Abs(float64(tail.row - head.row)) > 1.0 && math.Abs(float64(tail.col - head.col)) > 1.0 {
			moveKnot(tail, tail.row, tail.col + (head.col - tail.col) / 2)
			moveKnot(tail, tail.row + (head.row - tail.row) / 2, tail.col)
		} else if math.Abs(float64(tail.row - head.row)) > 1.0 {
			moveKnot(tail, tail.row, tail.col + int(float64(head.col - tail.col)))
			moveKnot(tail, tail.row + (head.row - tail.row) / 2, tail.col)
		} else {
			moveKnot(tail, tail.row, tail.col + (head.col - tail.col) / 2)
			moveKnot(tail, tail.row + int(float64(head.row - tail.row)), tail.col)
		}
	}
}

func moveHead(head *Knot, orientation string) {
	var drow, dcol int
	switch orientation {
	case "U":
		drow = -1 
	case "D":
		drow = 1
	case "L":
		dcol = -1
	case "R":
		dcol = 1
	}
	moveKnot(head, head.row + drow, head.col + dcol)
}

func moveKnot(head *Knot, newRow, newCol int) {
	head.row = newRow
	head.col = newCol
}

func printGrid(bridge [][]string, tails []Knot, head Knot) {
	for row, bridgeRow := range bridge {
		cycleCols:
		for col, cell := range bridgeRow {
			if row == head.row && col == head.col {
				fmt.Print(head.name)
			} else {
				for _, tail := range tails {
					if row == tail.row && col == tail.col {
						fmt.Print(tail.name)
						continue cycleCols
					} 
				}
				if cell == "" {
					fmt.Print("0")
				} else {
					fmt.Print(cell)
				}
			}
		}
		fmt.Println()
	}
	fmt.Println("--------------")
}

func getDirections() (dirs []Direction) {
	var orientation string
	var distance int
	scanner := bufio.NewScanner(os.Stdin) 
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(line, "%s %d", &orientation, &distance)
		dirs = append(dirs, Direction{orientation, distance})
	}
	return
} 

func getBridgeGrid(directions []Direction) (bridge [][]string) {
	var cols, rows int
	var startRow, startCol int
	rows, cols, startRow, startCol = getGridDimension(directions)
	fmt.Println(rows, cols, startRow, startCol)
	bridge = make([][]string, rows)
	for i, _ := range bridge {
		bridge[i] = make([]string, cols)
	}
	return 
}

//Returns Grid dimensions and starting row and col for the H not to go out of range
func getGridDimension(dirs []Direction) (int, int, int, int) {
	var maxRows, minRows, maxCols, minCols int
	var rows, cols int = 0, 0
	for _, dir := range dirs {
		switch dir.orientation {
		case "D":
			rows -= dir.distance
		case "L":
			cols -= dir.distance
		case "U":
			rows += dir.distance
		case "R":
			cols += dir.distance
		}
		if rows > maxRows {
			maxRows = rows
		}
		if rows < minRows {
			minRows = rows
		}
		if cols > maxCols {
			maxCols = cols
		}
		if cols < minCols {
			minCols = cols
		}
	}
	return maxRows - minRows + 1, maxCols - minCols + 1, maxRows, - minCols
}