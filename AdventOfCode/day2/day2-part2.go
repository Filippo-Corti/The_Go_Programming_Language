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
	case "A":
		return ROCK
	case "B":
		return PAPER
	case "C":
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
		totalScore += playRockPaperScissors(Shape(moves[0]), moves[1])
	}
	fmt.Println("Total Score:", totalScore)
}

//Given the Shape played by your opponent and the outcome of the match (in this order), returns your score for the round
func playRockPaperScissors(opponentShape Shape, outcome string) int {
	if outcome == "Y" { //DRAW
		return DRAW + opponentShape.Value()
	} 
	if outcome == "X" { //LOSS
		return LOSS + getLosingShape(opponentShape).Value()
	}
	// (myShape.Value() - opponentShape.Value() + 3) % 3 == 1 {
	return WIN + getWinningShape(opponentShape).Value()
}

//Returns the winning shape based on what the opponent shape is
func getWinningShape(opponent Shape) Shape {
	if opponent.Value() == ROCK {
		return Shape("B")
	}
	if opponent.Value() == PAPER {
		return Shape("C")
	}
	return Shape("A")
}

//Returns the losing shape based on what the opponent shape is
func getLosingShape(opponent Shape) Shape {
	if opponent.Value() == ROCK {
		return Shape("C")
	}
	if opponent.Value() == PAPER {
		return Shape("A")
	}
	return Shape("B")
}