package main

/* --- Day 14: Regolith Reservoir */

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Grid struct {
	grid [][]Cell
	fallingPoint Cell
}

type Cell struct {
	row, col int
	fill int //0 -> Air, 1 -> Rock, 2 -> Sand
}

type Point struct {
	row, col int
}

type Polyline []Point

const (
	AIR = 0
	ROCK = 1
	SAND = 2
)


func main() {
	polylines := parseInput()
	rocks := parsePolylinesToRocks(polylines)
	grid := generateGrid(rocks)
	fmt.Println(countFallingSand(grid))
	printGrid(grid)
}

//Simulates the fall of units of sand until one of the units falls off the grid
//It returns the number of fallen units
func countFallingSand(grid Grid) (c int) {
	rows := len(grid.grid)
	unitsFalling:
	for {
		//Sand starts at the top
		sand := grid.fallingPoint
		unitFalling:
		for {
			if sand.row + 1 >= rows {
				break unitsFalling
			}
			if isAir(grid.grid[sand.row + 1][sand.col]) {
				//Down is free
				sand.row++
			} else {
				if sand.col - 1 < 0 {
					break unitsFalling
				}
				if isAir(grid.grid[sand.row + 1][sand.col - 1])  {
					//Left is free
					sand.row++
					sand.col--
				} else if isAir(grid.grid[sand.row + 1][sand.col + 1]) {
					//Right is free
					sand.row++
					sand.col++
				} else {
					//Sand has finished falling
					break unitFalling
				}
			} 
		}
		c++
		grid.grid[sand.row][sand.col].fill = SAND
	}
	return c
}

func isAir(c Cell) bool {
	return c.fill == AIR
}

//Prints the grid in a nice way :)
func printGrid(grid Grid) {
	for _, gridRow := range grid.grid {
		for _, cell := range gridRow {
			if cell == grid.fallingPoint {
				fmt.Print("+")
				continue
			}
			if cell.fill == AIR {
				fmt.Print(".")
			}
			if cell.fill == ROCK {
				fmt.Print("#")
			}
			if cell.fill == SAND {
				fmt.Print("o")
			}
		}
		fmt.Println()
	}
}

//Receives the list of rocks inside the grid and generates a grid
//whose dimensions are optimized based on the min and max values of the rocks
func generateGrid(rocks []Point) (grid Grid) {
	var maxRow, minCol, maxCol int = 0, 999, 0
	//Create Grid
	for _, rock := range rocks {
		maxRow = max(rock.row, maxRow)
		minCol = min(rock.col, minCol)
		maxCol = max(rock.col, maxCol)
	}
	rows := maxRow + 1
	cols := maxCol - minCol + 1
	grid.fallingPoint = Cell{0, 500 - minCol, AIR}
	grid.grid = make([][]Cell, rows)
	for i := 0; i < rows; i++ {
		grid.grid[i] = make([]Cell, cols)
		for j := 0; j < cols; j++ {
			grid.grid[i][j] = Cell{i, j, AIR}
		}
	}
	//Fill the Grid
	for _, rock := range rocks {
		row := rock.row
		col := rock.col - minCol
		grid.grid[row][col] = Cell{row, col, ROCK}
	}
	return
}

//Receives a slice of polylines, draws the lines and returns every single point that was drawn
//during the process
func parsePolylinesToRocks(polylines []Polyline) (points []Point) {
	for _, polyline := range polylines {
		points = append(points, parsePolylineToRocks(polyline)...)
	}
	return
}

//Receives a polyline, draws the lines and returns every single point that was drawn
//during the process
func parsePolylineToRocks(polyline Polyline) (points []Point) {
	points = append(points, polyline[0])
	for i := 0; i < len(polyline) - 1; i++ {
		points = append(points, drawLine(polyline[i], polyline[i + 1])...)
	}
	return
}

//Receives the starting point and the finish point and returns 
//All the points included in the line EXCEPT the starting point (!)
func drawLine(from, to Point) (points []Point) {
	if from.row == to.row {
		//Horizontal line
		if from.col < to.col {
			//Going Left to Right
			for i := from.col + 1; i <= to.col; i++ {
				points = append(points, Point{from.row, i})
			}
		} else {
			//Going Right to Left 
			for i := from.col - 1; i >= to.col; i-- {
				points = append(points, Point{from.row, i})
			}
		}
	} else {
		//Vertical line
		if from.row < to.row {
			//Going Up to Down
			for i := from.row + 1; i <= to.row; i++ {
				points = append(points, Point{i, from.col})
			} 
		} else {
			//Going Down to Up
			for i := from.row - 1; i >= to.row; i-- {
				points = append(points, Point{i, from.col})
			}
		}
	}
	return
}

//It reads the input from standard input and returns the corresponding
//slice of Polylines
func parseInput() (polylines []Polyline) {
	var i int 
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := parseLine(line)
		polylines = append(polylines, make(Polyline, 0))
		polylines[i] = append(polylines[i], parsedLine...)
		i++
	}
	return
}

//It receives a string containing a line of the syntax "X1,Y1 -> X2,Y2 -> ... " and
//returns the corresponding Polyline (slice of points)
func parseLine(line string) (polyline Polyline) {
	var x, y int
	items := strings.Split(line, "->")
	for _, item := range items {
		fmt.Sscanf(strings.TrimSpace(item), "%d,%d", &x, &y)
		polyline = append(polyline, Point{y, x})
	}
	return
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}