package main

import (
  "fmt"
  "math"
)

func main() {
  var n int
  fmt.Print("Insert Number n: ")
  fmt.Scan(&n)
  fmt.Println("List of perfect numbers <=n: ")
  printPerfectNumbersUpToN(n)
}

func printPerfectNumbersUpToN(n int) {
  for i := 1; i <= n; i++ {
    if isItPerfect(i) {
      fmt.Println(i, "is perfect")
    }
  }
}

func isItPerfect(x int) bool {
  return areTheyFriends(x, x)
}

func areTheyFriends(x,y int) bool {
  return x == getSumOfDivisors(y) && y == getSumOfDivisors(x)
}

func getSumOfDivisors(n int) (s int) {
  s += 1
  for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
    if n % i == 0 {
      s += i
      if i != n / i {
        s += (n / i)
      }
    }
  }
  return s
}
