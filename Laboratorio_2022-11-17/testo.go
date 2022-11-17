package main

import (
   "fmt"
)

func main() {
   p1 := newInt1(2)
   p2 := newInt2(6)
   sum := *p1 + *p2
   fmt.Println("somma:", sum)
}

func newInt1(n int) *int {
   p := new(int)
   *p = n
   return p
}

func newInt2(n int) *int {
   var temp int
   temp = n
   return &temp
}