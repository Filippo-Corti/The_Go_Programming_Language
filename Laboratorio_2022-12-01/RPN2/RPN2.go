package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    var stack []float64
    for i := 1; i < len(os.Args); i++ {
        input := os.Args[i]
        inputNum, err := strconv.ParseFloat(input, 64)
        if err != nil {
            if len(stack) < 2 {
                fmt.Println("Not Enough Data")
                continue
            }
            n2 := Pop(&stack)
            n1 := Pop(&stack)
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
    }
    fmt.Println(stack)
}