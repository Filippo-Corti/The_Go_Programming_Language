package main

/* --- Day 15: Beacon Exclusion Zone --- */

import (
	"bufio"
	"fmt"
	"math"
	"golang.org/x/exp/maps"
	"os"
)

const SIZE = 4000000

type Coord struct {
	row, col int
}

type Report struct {
	sensor, beacon Coord
}

func main() {
	reports := parseInput()
	beacon := findBeacon(reports)
	fmt.Println(beacon)
	fmt.Println(getTuningFrequency(beacon))
}


func getTuningFrequency(beacon Coord) int {
	return beacon.col * 4000000 + beacon.row
}


func findBeacon(reports []Report) (beacon Coord) {
	var i int
	var noBeaconCoords map[Coord]bool
	//Find Row
	for i = 0; i <= SIZE; i++ {
		noBeaconCoords = getNoBeaconCoords(reports, i)
		if len(noBeaconCoords) != SIZE + 1 {
			break
		}
	}
	//Find Col
	for j := 0; j <= SIZE; j++ {
		_, ok := noBeaconCoords[Coord{i, j}]
		if !ok {
			return Coord{i, j}
		}
	}
	return
}


func merge(s []map[Coord]bool) (m map[Coord]bool) {
	m = make(map[Coord]bool)
	for _, report := range s {
		maps.Copy(m, report)
	}
	return
}

func getNoBeaconCoords(reports []Report, row int) map[Coord]bool {
	coordsSlice := make([]map[Coord]bool, len(reports))
	for i, report := range reports {
		coordsSlice[i] = make(map[Coord]bool)
		list := getNoBeaconCoordsOneReport(report, row)
		for _, el := range list {
			coordsSlice[i][el] = true
		}
	}
	return merge(coordsSlice)
}

func getNoBeaconCoordsOneReport(report Report, row int) (coords []Coord) {
	distFromBeacon := calcManhattanDistance(report.sensor, report.beacon)
	distFromRow := int(math.Abs(float64(report.sensor.row - row)))
	for i := report.sensor.col - (distFromBeacon - distFromRow); i <= report.sensor.col + (distFromBeacon - distFromRow); i++ {
		if i >= 0 && i <= SIZE {
			coords = append(coords, Coord{row, i})
		}
	}
	return
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