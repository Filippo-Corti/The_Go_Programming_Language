package main

/* --- Day 8: Treetop Tree House --- */

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	treeGrid := getTreeGrid()
	fmt.Println(countVisibleTrees(getVisibleTrees(treeGrid)))
}

func getTreeGrid() (grid [][]int) {
	grid = make([][]int, 0)
	var i int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []int{})
		grid[i] = append(grid[i], getTreeRow(line)...)
		i++
	}
	return
}

func getTreeRow(s string) (row []int) {
	for _, r := range s {
		row = append(row, int(r) - int('0'))
	}
	return
}

//Returns a grid of the same size of treeGrid, in which visible trees have value -1
func getVisibleTrees(treeGrid [][]int) (visibleTrees [][]int) {
	visibleTrees = make([][]int, len(treeGrid))
	for i, _ := range visibleTrees {
		visibleTrees[i] = make([]int, len(treeGrid[i]))
	}
	getVisibleTreesFromTop(treeGrid, visibleTrees)
	getVisibleTreesFromBottom(treeGrid, visibleTrees)
	getVisibleTreesFromLeft(treeGrid, visibleTrees)
	getVisibleTreesFromRight(treeGrid, visibleTrees)
	return
}

func getVisibleTreesFromTop(treeGrid, visibleTrees [][]int) {
	var minHeight int = -1
	for col := 0; col < len(treeGrid[0]); col++ {
		for row := 0; row < len(treeGrid); row++ {
			if treeGrid[row][col] > minHeight {
				if visibleTrees[row][col] != -1 {
					visibleTrees[row][col] = -1
				}
				minHeight = treeGrid[row][col]
			}
		}
		minHeight = -1
	}
}

func getVisibleTreesFromBottom(treeGrid, visibleTrees [][]int) {
	var minHeight int = -1
	for col := len(treeGrid[0]) - 1; col >= 0; col-- {
		for row := len(treeGrid) - 1; row >= 0; row-- {
			if treeGrid[row][col] > minHeight {
				if visibleTrees[row][col] != -1 {
					visibleTrees[row][col] = -1
				}
				minHeight = treeGrid[row][col]
			}
		}
		minHeight = -1
	}
}

func getVisibleTreesFromLeft(treeGrid, visibleTrees [][]int) {
	var minHeight int = -1
	for row := 0; row < len(treeGrid[0]); row++ {
		for col := 0; col < len(treeGrid); col++ {
			if treeGrid[row][col] > minHeight {
				if visibleTrees[row][col] != -1 {
					visibleTrees[row][col] = -1
				}
				minHeight = treeGrid[row][col]
			}
		}
		minHeight = -1
	}
}

func getVisibleTreesFromRight(treeGrid, visibleTrees [][]int) {
	var minHeight int = -1
	for row := len(treeGrid[0]) - 1; row >= 0; row-- {
		for col := len(treeGrid) - 1; col >= 0; col-- {
			if treeGrid[row][col] > minHeight {
				if visibleTrees[row][col] != -1 {
					visibleTrees[row][col] = -1
				}
				minHeight = treeGrid[row][col]
			}
		}
		minHeight = -1
	}
}

//Counts values of visibleTrees equal to -1
func countVisibleTrees(visibleTrees [][]int) (c int) {
	for i := 0; i < len(visibleTrees); i++ {
		for _, v := range visibleTrees[i] {
			if v == -1 {
				c++
			}
		}
	}
	return
}