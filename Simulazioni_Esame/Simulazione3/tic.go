package main

import (
	"fmt"
	"os"
	"strconv"
)

type Orario struct {
	h, m, s int
}

func (o Orario) String() string {
	return fmt.Sprintf("%d:%d:%d", o.s, o.m, o.h)
}

func newOrario(s, m, h int) (bool, Orario) {
	if s < 0 || s > 59 || m < 0 || m > 59 || h < 0 || h > 23 {
		return false, Orario{0,0,0}
	}
	return true, Orario{h, m, s}
}

func tic(orario *Orario) {
	orario.s += 1
	if orario.s > 59 {
		orario.m += 1
		orario.s = 0
		if orario.m > 59 {
			orario.h += 1
			orario.m = 0
			if orario.h > 23 {
				orario.h = 0
			}
		}
	}
}

func main() {
	s, _ := strconv.Atoi(os.Args[1])
	m, _ := strconv.Atoi(os.Args[2])
	h, _ := strconv.Atoi(os.Args[3])
	time, _ := strconv.Atoi(os.Args[4])
	_ = time
	ok, orario := newOrario(s, m, h)
	if !ok {
		fmt.Println("parametri non validi")
		os.Exit(0)
	}
	fmt.Println(orario)
	for i := 0; i < time; i++ {
		tic(&orario)
	}
	fmt.Println(orario)
}