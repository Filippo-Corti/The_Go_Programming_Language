package main

import (
    "fmt"
    "strings"
    "strconv"
    "os"
)

type Rettangolo struct {
    base, altezza int
}

func (r *Rettangolo) String() (s string) {
    return strings.Repeat(strings.Repeat(".", r.base) + "\n", r.altezza)
}

func main() {
    base, _ := strconv.Atoi(os.Args[1])
    altezza, _ := strconv.Atoi(os.Args[2])
    rett := Rettangolo{base: base, altezza: altezza}
    fmt.Println(rett.String())
}