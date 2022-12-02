package main

import (
)

func Push(stack *[]float64, n float64) {
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