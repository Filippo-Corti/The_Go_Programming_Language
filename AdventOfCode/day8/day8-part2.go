package main

/* --- Day 8: Treetop Tree House --- */

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	treeGrid := getTreeGrid()
	//fmt.Println(getScenicScore(treeGrid, 1, 1))
	fmt.Println(getBestScenicScore(treeGrid))
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

func getBestScenicScore(treeGrid [][]int) (scenicScore int) {
	for row := 0; row < len(treeGrid[0]); row++ {
		for col := 0; col < len(treeGrid); col++ {
			temp := getScenicScore(treeGrid, row, col)
			if temp != 0 {
				fmt.Println(treeGrid[row][col], row, col, getScenicScore(treeGrid, row, col))
			}
			if temp > scenicScore {
				scenicScore = temp
			}
		}
	}
	return
}

//Returns the number of trees visible from treeGrid[treeRow][treeCol]
func getScenicScore(treeGrid [][]int, treeRow, treeCol int) (scenicScore int) {
	scenicScore = getScenicScoreFromTop(treeGrid, treeRow, treeCol)
	//fmt.Println(getScenicScoreFromTop(treeGrid, treeRow, treeCol))
	scenicScore *= getScenicScoreFromBottom(treeGrid, treeRow, treeCol)
	//fmt.Println(getScenicScoreFromBottom(treeGrid, treeRow, treeCol))
	scenicScore *= getScenicScoreFromLeft(treeGrid, treeRow, treeCol)
	//fmt.Println(getScenicScoreFromLeft(treeGrid, treeRow, treeCol))
	scenicScore *= getScenicScoreFromRight(treeGrid, treeRow, treeCol)
	//fmt.Println(getScenicScoreFromRight(treeGrid, treeRow, treeCol))
	return 
}

func getScenicScoreFromTop(treeGrid [][]int, treeRow, treeCol int) (scenicScore int) {
	var row int
	for row = treeRow - 1; row >= 0 && treeGrid[row][treeCol] < treeGrid[treeRow][treeCol]; row-- {
		scenicScore++
	}
	if row >= 0 && treeGrid[row][treeCol] >= treeGrid[treeRow][treeCol] {
		scenicScore++
	}
	return
}

func getScenicScoreFromBottom(treeGrid [][]int, treeRow, treeCol int) (scenicScore int) {
	var row int
	for row = treeRow + 1; row < len(treeGrid) && treeGrid[row][treeCol] < treeGrid[treeRow][treeCol]; row++ {
		scenicScore++
	}
	if row < len(treeGrid) && treeGrid[row][treeCol] >= treeGrid[treeRow][treeCol] {
		scenicScore++
	}
	return
}

func getScenicScoreFromLeft(treeGrid [][]int, treeRow, treeCol int) (scenicScore int) {
	var col int
	for col = treeCol - 1; col >= 0 && treeGrid[treeRow][col] < treeGrid[treeRow][treeCol]; col-- {
		scenicScore++
	}
	if col >= 0 && treeGrid[treeRow][col] >= treeGrid[treeRow][treeCol] {
		scenicScore++
	}
	return
}

func getScenicScoreFromRight(treeGrid [][]int, treeRow, treeCol int) (scenicScore int) {
	var col int
	for col = treeCol + 1; col < len(treeGrid[treeRow]) && treeGrid[treeRow][col] < treeGrid[treeRow][treeCol]; col++ {
		scenicScore++
	}
	if col < len(treeGrid[treeRow]) && treeGrid[treeRow][col] >= treeGrid[treeRow][treeCol] {
		scenicScore++
	}
	return
}