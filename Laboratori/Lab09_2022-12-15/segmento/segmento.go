package main

import (
	"fmt"
	"strings"
)

type Segmento struct {
	estremi byte
	interno byte
	orizzontale bool
	lunghezza int
}

func (s *Segmento) String() string {
	if s.orizzontale {
		return string(s.estremi) +  strings.Repeat(string(s.interno), s.lunghezza - 2) + string(s.estremi)
	} else {
		return string(s.estremi) + "\n" +  strings.Repeat(string(s.interno) + "\n", s.lunghezza - 2) + string(s.estremi)
	}
}

func (s *Segmento) Allunga(n int) {
	s.lunghezza += n
}

func main() {
	var estremi, interno string
	var orizzontale bool
	var lunghezza int
	fmt.Scan(&estremi, &interno, &orizzontale, &lunghezza)
	seg := Segmento{estremi[0], interno[0], orizzontale, lunghezza} 
	fmt.Println(seg.String())
	seg.Allunga(5)
	fmt.Println(seg.String())
}