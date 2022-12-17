package main

/*
	Rock Paper Scissors following a strategy guide
*/

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Shape string

func (s Shape) Value() int {
	switch s {
	case "X", "A":
		return ROCK
	case "Y", "B":
		return PAPER
	case "Z", "C":
		return SCISSOR
	default:
		return -1
	}
}

const (
	ROCK = 1
	PAPER = 2
	SCISSOR = 3
	LOSS = 0
	DRAW = 3
	WIN = 6
)

func main() {
	var totalScore int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		moves := strings.Split(line, " ")
		totalScore += playRockPaperScissors(Shape(moves[0]), Shape(moves[1]))
	}
	fmt.Println("Total Score:", totalScore)
}

//Given the Shape played by your opponent and you (in this order), returns your score for the round
func playRockPaperScissors(opponentShape, myShape Shape) int {
	if opponentShape.Value() == myShape.Value() {
		return DRAW + myShape.Value() //DRAW
	}
	if (myShape.Value() - opponentShape.Value() + 3) % 3 == 1 {
		return WIN + myShape.Value()
	}
	return LOSS + myShape.Value()
}