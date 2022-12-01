package main

import (
    "fmt"
    "strconv"
)

func main() {
    var input string
    var stack []float64
    for {
        fmt.Print("Next? (+, -, *, /, q o un numero) ")
        fmt.Scan(&input)
        if input == "q" {
            break
        }
        inputNum, err := strconv.ParseFloat(input, 64)
        if err != nil {
            if len(stack) < 2 {
                fmt.Println("Not Enough Data")
                continue
            }
            n1 := Pop(&stack)
            n2 := Pop(&stack)
            switch input {
            case "+":
                Push(&stack, n1 + n2)
            case "-":
                Push(&stack, n1 - n2)
            case "*":
                Push(&stack, n1 * n2)
            case "/":
                Push(&stack, n1 / n2)
            }
        } else {
            Push(&stack, inputNum)
        }
        fmt.Println(stack)
    }
}