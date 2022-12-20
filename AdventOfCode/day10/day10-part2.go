package main

/* --- Day 10: Cathode-Ray Tube --- */

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type CPU struct {
	cycle int
	x int
	signalStrenghts int
}

func (cpu *CPU) Noop() {
	if (cpu.cycle % 40) >= cpu.x - 1 && (cpu.cycle % 40) <= cpu.x + 1 {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}
	cpu.cycle++
	if cpu.cycle % 40 == 0 {
		fmt.Println()
	}
	if (cpu.cycle - 20) % 40 == 0 {
		cpu.signalStrenghts += cpu.cycle * cpu.x
	}
}

func (cpu *CPU) AddX(x int) {
	for i := 0; i < 2; i++ {
		cpu.Noop()
		if i == 1 {
			cpu.x += x
		}
	}
}

func main() {
	cpu := CPU{0,1, 0}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		splittedLine := strings.Fields(line)
		if len(splittedLine) == 1 {
			cpu.Noop()
		} else {
			x, _ := strconv.Atoi(splittedLine[1])
			cpu.AddX(x)
		}
	}
	fmt.Println()
	fmt.Println(cpu.signalStrenghts)
}