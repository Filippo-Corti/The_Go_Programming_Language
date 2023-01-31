package main

import "fmt"

func main() {
  var n int
  fmt.Print("Inserisci un numero: ")
  fmt.Scan(&n)
  for i := 1; i <= n; i++ {
    printCollatzLength(i)
  }
}

func nextCollatz(n int) int {
  if n % 2 == 0 {
    return n / 2
  } else {
    return n * 3 + 1
  }
}

func printCollatzLength(n int) {
  fmt.Print(n, "\t")
  for {
    fmt.Print("*")
    if n == 1 {
      break
    }
    n = nextCollatz(n)
  }
  fmt.Println()
}
