package main

/* --- Day 13: Distress Signal --- */

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
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
	sort.Slice(input, func(i, j int) bool {
		v := isPairCorrect(Pair{input[i], input[j]})
		return v == 0
	})
	printInput(input)
	fmt.Println(getDecoderKey(input))
}

func comparePackets(p1, p2 Packet) bool {
	return fmt.Sprintf("%v", p1) == fmt.Sprintf("%v", p2)
}

//It returns the position of the two divider packets [[2]] and [[6]] multiplied together
func getDecoderKey(input []Packet) int {
	var i, j int
	divider1 := parsePacket("[[2]]")
	divider2 := parsePacket("[[6]]")
	for x, p := range input {
		if comparePackets(p, divider1) {
			i = x + 1
			continue
		}
		if comparePackets(p, divider2) {
			j = x + 1
		}
	}
	return i * j
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
func printInput(input []Packet) {
	for _, p := range input {
		fmt.Println(p)
	}
}

//Return the list of Packets contained in the input day13.txt
func parseInput() (packets []Packet) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){ //First Packet
		line := scanner.Text()
		if len(line) == 0 {
			//Ignore empty lines
			continue
		}
		packet := parsePacket(line)
		packets = append(packets, packet)
	}
	packets = append(packets, parsePacket("[[2]]"))
	packets = append(packets, parsePacket("[[6]]"))
	return
}

//Returns the corresponding Packet to the string s received as parameter
func parsePacket(s string) (packet Packet) {
	json.Unmarshal([]byte(s), &packet)
	convertToInts(&packet)
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