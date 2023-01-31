package main

import "fmt"

func main() {
  rombo(19)
}

func rombo(n int) {
  triangolo5(n)
  triangolo6(n)
}

/*
  *
 ***
*****
*/
func triangolo5(n int) {
  rows := (n / 2) + (n % 2)
  for r := 0; r < rows; r++ {
    //Spazi prima
    for a := 0; a < rows - r - 1; a++ {
      fmt.Print(" ")
    }
    //Asterischi
    for a := 0; a < r * 2 + 2 - (n % 2); a++ {
      fmt.Print("*")
    }
    fmt.Println(" ")
  }
}
/*
*****
 ***
  *
*/
func triangolo6(n int) {
  rows := (n / 2) + (n % 2)
  for r := 0; r < rows; r++ {
    //Spazi prima
    for a := 0; a < r; a++ {
      fmt.Print(" ")
    }
    //Asterischi
    for a := 0; a < n - r * 2; a++ {
      fmt.Print("*")
    }
    fmt.Println(" ")
  }
}

/*
*****
*****
*****
*****
*****
*/
func quadrato(n int) {
  fmt.Println("---------Quadrato---------")
  for r := 0; r < n; r++ {
    for c := 0; c < n; c++ {
      fmt.Print("*")
    }
    fmt.Println()
  }
}


/*
*
**
***
****
*****
*/
func triangolo1(n int) {
  fmt.Println("---------Triangolo 1---------")
  for r := 0; r < n; r++ {
    for c := 0; c < r + 1; c++ {
      fmt.Print("*")
    }
    fmt.Println()
  }
}

/*
*****
****
***
**
*
*/
func triangolo2(n int) {
  fmt.Println("---------Triangolo 2---------")
  for r := 0; r < n; r++ {
    for c := 0; c < n - (r + 1); c++ {
      fmt.Print("*")
    }
    fmt.Println()
  }
}

/*
    *
   **
  ***
 ****
*****
*/
func triangolo3(n int) {
  fmt.Println("---------Triangolo 3---------")
  for r := 0; r < n; r++ {
    for b := 0; b < n - (r + 1); b++ {
        fmt.Print(" ")
    }
    for c := 0; c < r + 1; c++ {
      fmt.Print("*")
    }
    fmt.Println()
  }
}

/*
*****
 ****
  ***
   **
    *
*/
func triangolo4(n int) {
  fmt.Println("---------Triangolo 4---------")
  for r := 0; r < n; r++ {
    for b := 0; b < r; b++ {
        fmt.Print(" ")
    }
    for c := 0; c < n - r; c++ {
      fmt.Print("*")
    }
    fmt.Println()
  }
}
