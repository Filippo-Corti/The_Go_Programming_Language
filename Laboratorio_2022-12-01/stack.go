package main

import (
    "fmt"
)

func main() {
    var stack []float64
    _ = stack
    var scelta string
    runningLoop:
    for {
        fmt.Println("Operazione (push/pop/top/empty/quit)?")
        fmt.Scan(&scelta)
        switch scelta {
        case "push":
            Push(&stack)
        case "pop":
            fmt.Println("Popped:", Pop(&stack))
        case "top":
            fmt.Println("On Top:", Top(&stack))
        case "empty":
            fmt.Println(Empty(&stack))
        case "quit":
            break runningLoop
        }
        fmt.Println(stack)
    }
}

func Push(stack *[]float64) {
    var n float64
    fmt.Println("Valore?")
    fmt.Scan(&n)
    (*stack) = append([]float64{n}, (*stack)...) 
}

func Pop(stack *[]float64) float64{
    head := (*stack)[0]
    (*stack) = (*stack)[1:]
    return head
}

func Top(stack *[]float64) float64{
    return (*stack)[0]
}

func Empty(stack *[]float64) bool{
    return len((*stack)) == 0
}