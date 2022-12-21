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
    grid, _, endPoint := getGrid()
    fmt.Println(getBestSpot(grid, endPoint))
    //fmt.Println(dijkstra(grid, endPoint, startPoint))
}

func getBestSpot(grid [][]Square, end Square) int {
    min := 99999
    for _, gridRow := range grid {
        for _, square := range gridRow {
            if square.value == 'a' && isItNextTo(grid, square, 'b')  {
                val := dijkstra(grid, end, square)
                if val < min {
                    min = val
                }
                fmt.Println(val)
            }
        }
    }
    return min
}

func isItNextTo(grid [][]Square, square Square, r rune) bool {
    cols := len(grid[0])
    rows := len(grid)
    if square.col + 1 < cols {
        if grid[square.row][square.col + 1].value == r {
            return true
        }
    }
    if square.col - 1 >= 0 {
        if grid[square.row][square.col - 1].value == r {
            return true
        }
    }
    if square.row + 1 < rows {
        if grid[square.row + 1][square.col].value == r {
            return true
        }
    }
    if square.row - 1 >= 0 {
        if grid[square.row - 1][square.col].value == r {
            return true
        }
    }
    return false
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
        sa := &grid[lowestValued.row][lowestValued.col]
        // - Right
        if lowestValued.col + 1 < cols {
            sr := &grid[lowestValued.row][lowestValued.col + 1]
            if canIMoveTo(*sa, *sr) && dijkstraTable[*sa] + 1 < dijkstraTable[*sr] {
                dijkstraTable[*sr] = dijkstraTable[*sa] + 1
            }
        }
        // - Left
        if lowestValued.col - 1 >= 0 {
            sl := &grid[lowestValued.row][lowestValued.col - 1]
            if canIMoveTo(*sa, *sl)  && dijkstraTable[*sa] + 1 < dijkstraTable[*sl] {
                dijkstraTable[*sl] = dijkstraTable[*sa] + 1
            }
        }
        // - Top
        if lowestValued.row - 1 >= 0 {
            st := &grid[lowestValued.row - 1][lowestValued.col]
            if canIMoveTo(*sa, *st)  && dijkstraTable[*sa] + 1 < dijkstraTable[*st] {
                dijkstraTable[*st] = dijkstraTable[*sa] + 1
            }
        }
        // - Bottom
        if lowestValued.row + 1 < rows {
            sb := &grid[lowestValued.row + 1][lowestValued.col]
            if canIMoveTo(*sa, *sb) && dijkstraTable[*sa] + 1 < dijkstraTable[*sb] {
                dijkstraTable[*sb] = dijkstraTable[*sa] + 1
            }
        }
        //Remove it from the unvisited list
        unvisited = remove(unvisited, lowestValued) 
    }
    return dijkstraTable[grid[end.row][end.col]]
}

//Checks if two squares are adjacent
func canIMoveTo(moveFrom, moveTo Square) bool {
    return int(moveTo.value) >= int(moveFrom.value) - 1
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
    f, _ := os.Open("day12.txt")
    defer f.Close()
    var row, col int
    scanner := bufio.NewScanner(f)
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