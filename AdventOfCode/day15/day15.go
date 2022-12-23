package main

/* --- Day 15: Beacon Exclusion Zone --- */

import (
	"bufio"
	"fmt"
	"math"
	"os"
)


type Coord struct {
	row, col int
}

type Report struct {
	sensor, beacon Coord
}

func main() {
	reports := parseInput()
	noBeaconCoords := getNoBeaconCoords(reports, 2000000)
	fmt.Println(removeBeaconsAndCount(reports, noBeaconCoords))
}

func removeBeaconsAndCount(reports []Report, coords map[Coord]bool) int {
	//Remove Beacons
	for _, report := range reports {
		delete(coords, report.beacon)
	}

	return len(coords)
}

func getNoBeaconCoords(reports []Report, row int) (coords map[Coord]bool) {
	coords = make(map[Coord]bool)
	for _, report := range reports {
		list := getNoBeaconCoordsOneReport(report, row)
		for _, el := range list {
			coords[el] = true
		}
	}
	return
}

func getNoBeaconCoordsOneReport(report Report, row int) (coords []Coord) {
	distFromBeacon := calcManhattanDistance(report.sensor, report.beacon)
	distFromRow := int(math.Abs(float64(report.sensor.row - row)))
	for i := report.sensor.col - (distFromBeacon - distFromRow); i <= report.sensor.col + (distFromBeacon - distFromRow); i++ {
		coords = append(coords, Coord{row, i})
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