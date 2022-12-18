package main

/* --- Day 4: Camp Cleanup --- */

import (
	"fmt"
	"io"
)

type Interval struct {
	start, end int
}

func main() {
	var start1, end1, start2, end2, fullyContainedCounter int
	for {
		_, err := fmt.Scanf("%d-%d,%d-%d", &start1, &end1, &start2, &end2)
		if err == io.EOF {
			break
		}
		interval1 := Interval{start1, end1}
		interval2 := Interval{start2, end2}
		if doTheyContainEachOther(interval1, interval2) {
			fullyContainedCounter++
		}
	}
	fmt.Println("Fully Contained Intervals: ", fullyContainedCounter)
}

func doTheyContainEachOther(i1, i2 Interval) bool {
	return doesItOverlap(i1, i2) || doesItOverlap(i2, i1)
}

func doesItOverlap(i1, i2 Interval) bool {
	return i1.start >= i2.start && i1.start <= i2.end
}