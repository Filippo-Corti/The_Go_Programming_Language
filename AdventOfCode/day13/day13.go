package main

/* --- Day 13: Distress Signal --- */

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

const (
	FLOAT = "float64"
	INTEGER = "int"
	LIST = "[]interface {}"
)

type Packet []any

type Pair struct {
	left, right Packet
}

func getType(a any) string {
	return fmt.Sprintf("%T", a)
}

func main() {
	input := parseInput()
	fmt.Println(sumIndicesOfCorrectPairs(input))
}

//Counts the number of pairs for which isPairCorrect() return true
func sumIndicesOfCorrectPairs(pairs []Pair) (c int) {
	for i, pair := range pairs {
		fmt.Println("==== Pair", i + 1, "====")
		if isPairCorrect(pair) == 0 {
			fmt.Println("Correct", i+1)
			c+=i+1
		}
	}
	return 
}

//Returns whether the pair is (0), isn't in the correct order (1) or it can't tell (2)
func isPairCorrect(pair Pair) int {
	fmt.Println("-Compare", pair.left, "vs", pair.right)
	for i := 0;; i++ {
		if i >= len(pair.left) && i >= len(pair.right) {
			return 2 //It can't tell cause they are the same length
		}
		//I first check if any of the lists ran out of items 
		if i >= len(pair.left) {
			fmt.Println("Left side ran out of items first so they are in the RIGHT order", pair.left)
			return 0
		}
		if i >= len(pair.right) {
			fmt.Println("Left side ran out of items first so they are in the WRONG order", pair.right)
			return 1
		}
		left := pair.left[i]
		right := pair.right[i]
		fmt.Println("--Compare", left, "vs", right)
		if getType(left) == getType(right) {
			if getType(left) == INTEGER {
				//If they are both Integers, I need to compare them 
				if left.(int) == right.(int) {
					continue
				}
				if left.(int) < right.(int) {
					fmt.Println("Left side is smaller, so inputs are in the right order")
					return 0
				} else {
					fmt.Println("Right side is smaller, so inputs are in the wrong order")
					return 1
				}
			} else {
				//If they are both Lists, I need to compare every item inside 
				res := isPairCorrect(Pair{Packet(left.([]interface {})), Packet(right.([]interface {}))})
				if res != 2 {
					return res
				}
			}
		} else {
			var res int
			//If they are of different types, I need to convert the int to a list and retry
			if getType(left) == INTEGER {
				fmt.Println("Mixed types; convert left to list and retry comparison")
				res = isPairCorrect(Pair{[]any{left}, Packet(right.([]interface {}))})
			} else {
				fmt.Println("Mixed types; convert right to list and retry comparison")
				res = isPairCorrect(Pair{Packet(left.([]interface {})), []any{right}})
			}
			if res != 2 {
				return res
			}
		}
	}
	return 2
}

//Prints a slice of Pair in a more readable way
func printInput(input []Pair) {
	for _, pair := range input {
		fmt.Println(pair)
	}
}

//Return the list of Pair contained in the input day13.txt
func parseInput() (pairs []Pair) {
	scanner := bufio.NewScanner(os.Stdin)
	for { //First Packet
		pair, done := parsePair(scanner)
		pairs = append(pairs, pair)
		if done {
			return
		}
	}
	return
}

//Recevies the scanner to read the input from and returns a single Pair (couple of []any) and a flag which 
//is true if the file is finished or false if there is more to read
func parsePair(scanner *bufio.Scanner) (pair Pair, done bool) {
	var packet1, packet2 Packet
	scanner.Scan()
	line := scanner.Text()
	json.Unmarshal([]byte(line), &packet1)
	convertToInts(&packet1)
	pair.left = packet1
	scanner.Scan()
	line = scanner.Text()
	json.Unmarshal([]byte(line), &packet2)
	convertToInts(&packet2)
	pair.right = packet2
	//Skip next empty line
	ok := scanner.Scan()
	if !ok {
		done = true
		return
	}
	return
}

//Converts any element of type float64 inside the slice of any into an integer 
func convertToInts(s *Packet) {
	for i, el := range *s {
		if getType(el) == FLOAT {
			(*s)[i] = int(el.(float64))
		} else if getType(el) == LIST {
			l := Packet(((*s)[i]).([]interface {}))
			convertToInts(&l)
		}
	}
}