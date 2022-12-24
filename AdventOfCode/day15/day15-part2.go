package main

/* --- Day 15: Beacon Exclusion Zone --- */

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const SIZE = 4000000

type Coord struct {
	row, col int
}

type Interval struct {
	start, end int
}

type Intervals []Interval

type Report struct {
	sensor, beacon Coord
}

func main() {
	reports := parseInput()
	//fmt.Println(mergeIntervalIntoSlice([]Interval{Interval{0,10}, Interval{12,20}}, Interval{0, 14}))
	beacon := findBeacon(reports)
	fmt.Println(beacon)
	fmt.Println(getTuningFrequency(beacon))
}

func findBeacon(reports []Report) (beacon Coord) {
	intervals := map[int]Intervals{}
	//For each rombus
	for _, report := range reports {
		//For each row 
		dist := calcManhattanDistance(report.sensor, report.beacon)
		for row, i := report.sensor.row - dist, 0; row <= report.sensor.row; row, i = row + 1, i + 1{
			if row < 0 || row > SIZE {
				continue
			}
			interval := Interval{report.sensor.col - i, report.sensor.col + i}
			interval.start = max(interval.start, 0)
			interval.end = min(interval.end, SIZE)
			intervals[row] = mergeIntervalIntoSlice(intervals[row], interval)
		}
		for row, i := report.sensor.row + 1, dist - 1; row <= report.sensor.row + dist; row, i = row + 1, i - 1 {
			if row < 0 || row > SIZE {
				continue
			}
			interval := Interval{report.sensor.col - i, report.sensor.col + i}
			interval.start = max(interval.start, 0)
			interval.end = min(interval.end, SIZE)
			intervals[row] = mergeIntervalIntoSlice(intervals[row], interval)
		}
	}
	for i := 0; i <= SIZE; i++ {
		length := len(intervals[i])
		if length == 2 {
			fmt.Println(i, intervals[i])
			return Coord{i, intervals[i][0].end + 1}
		}
	}
	return Coord{-1, -1}
}

func insertInOrder(intervals []Interval, newInterval Interval) []Interval  {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}
	for i, interval := range intervals {
		if newInterval.start < interval.start {
			return append(intervals[:i], append([]Interval{newInterval}, intervals[i:]...)...)
		}
	}
	return append(intervals, newInterval)
}

func mergeIntervalIntoSlice(intervals []Interval, newInterval Interval) []Interval {
	intervals = insertInOrder(intervals, newInterval)
	for i := len(intervals) - 1; i > 0; i-- {
		newInt, ok := mergeIntervals(intervals[i], intervals[i - 1])
		if ok {
			intervals = append(intervals[: i - 1], append([]Interval{newInt}, intervals[i + 1 :]...)...)
		}
	}
	if len(intervals) == 2 {
		lastAttempt, ok := mergeIntervals(intervals[0], intervals[1])
		if ok {
			return []Interval{lastAttempt}
		}
	}
	return intervals
}

func mergeIntervals(interval1, interval2 Interval) (Interval, bool) {
	if interval1.start > interval2.start {
		interval1, interval2 = interval2, interval1
	}
	if interval1.start <= interval2.start && interval1.end >= interval2.start - 1{
		if interval1.end > interval2.end {
			return interval1, true
		} else {
			return Interval{interval1.start, interval2.end}, true
		}
	}
	return Interval{-1, -1}, false
}


func getTuningFrequency(beacon Coord) int {
	return beacon.col * 4000000 + beacon.row
}

func calcManhattanDistance(p1, p2 Coord) int {
	return int(math.Abs(float64(p1.col - p2.col)) + math.Abs(float64(p1.row - p2.row)))
}

func parseInput() (reports []Report) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		reports = append(reports, parseLine(line))
	}
	return
}

func parseLine(s string) (r Report) {
	var sensorRow, sensorCol, beaconRow, beaconCol int
	fmt.Sscanf(s, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorCol, &sensorRow, &beaconCol, &beaconRow)
	r.sensor = Coord{sensorRow, sensorCol}
	r.beacon = Coord{beaconRow, beaconCol}
	return
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}