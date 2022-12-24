package main

/* */

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Valve struct {
	name string 
	rate int
}

type Pair struct {
	v1, v2 Valve
}

const (
	MINUTES = 30
)

func main() {
	valves, links := parseInput()
	// dijkstraTables := calcDijkstraTables(valves, links)
	// //printTables(dijkstraTables)
	// fmt.Println(startMoving(dijkstraTables, valves))
	pairs := findPairs(valves, links)
	fmt.Println(pairs)
}

func createPath(pairs, valves) {
	
}

func findPairs(valves map[string]Valve, links map[Valve][]string) (pairs map[Pair]bool) {
	pairs = make(map[Pair]bool)
	for _, valve := range valves {
		neighbours := links[valve]
		for _, neighbour := range neighbours {
			pair := Pair{valve, valves[neighbour]}
			inverted := Pair{valves[neighbour], valve}
			_, ok := pairs[inverted]
			if !ok {
				pairs[pair] = true
			}
		}
	} 
	return
}









//For each minute I figure out the most efficient move and I do execute it
//It returns the total flow 
func startMoving(dijkstraTables map[Valve]map[Valve]int, valves map[string]Valve) (totalFlow int) {
	flowPerMin := 0
	currentValve := valves["AA"]
	openValves := []string{"AA"}
	for currentMinute := 1; currentMinute <= MINUTES; {
		//Find the best move
		nextValve, timeToGetThere := findBestMove(dijkstraTables[currentValve], currentMinute, openValves)
		fmt.Println("Moving to", nextValve, "taking", timeToGetThere, "minutes")
		//Update flowPerMin and totalFlow
		openValves = append(openValves, nextValve.name)
		totalFlow += flowPerMin * timeToGetThere
		flowPerMin += nextValve.rate
		//Update currentMinute
		if timeToGetThere == 0 {
			return totalFlow + flowPerMin * (MINUTES - currentMinute)
		}
		currentMinute += timeToGetThere + 1
		fmt.Println("Current Minute:", currentMinute)
		currentValve = nextValve
	}
	return
}

func containsString(list []string, str string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

func findBestMove(dijkstraTable map[Valve]int, currentMinute int, openValves []string) (Valve, int) {
	bestResult := 0
	bestValve := Valve{}
	bestMinute := 0
	for valve, minutesToGetTo := range dijkstraTable {
		if containsString(openValves, valve.name) || valve.rate == 0 {
			//If a valve is open I am not going to get there for nothing
			continue 
		}
		//the score consist in the amount of flow that would go through that valve 
		//from the moment I get to open it to the end of the time 
		score := (MINUTES - (currentMinute + minutesToGetTo)) * valve.rate
		if score > bestResult {
			bestResult = score
			bestMinute = minutesToGetTo
			bestValve = valve
		}
	}
	return bestValve, bestMinute
}

func printTables(dijkstraTables map[Valve]map[Valve]int) {
	for valve, dijkstraTable := range dijkstraTables {
		fmt.Println(valve, "\t", dijkstraTable)
	}
}
//Returns a map of maps
//Each map is the dijkstraTable of the key Valve and contains the path cost to move 
//to any Valve in the graph
func calcDijkstraTables(valves map[string]Valve, links map[Valve][]string) (dijkstraTables map[Valve]map[Valve]int) {
	dijkstraTables = make(map[Valve]map[Valve]int)
	for _, valve := range valves {
		dijkstraTables[valve] = calcDijkstraTable(valve, valves, links)
	}
	return
}

func calcDijkstraTable(valve Valve, valves map[string]Valve, links map[Valve][]string) (dijkstraTable map[Valve]int) {
	var visited, unvisited []Valve
	//Make all Valves unvisited 
	visited = make([]Valve, 0)
	unvisited = listItems(valves)
	//Set distance to infinity
	dijkstraTable = initTable(unvisited)
	//Set distance for the current Valve to zero
	dijkstraTable[valve] = 0
	//Loop through every Valve
	for len(unvisited) > 0 {
		//Find Valve with lowest value
		lowestValued := getLowestValued(dijkstraTable, visited)
		//Mark it as visited
		visited = append(visited, lowestValued)
		//Check if the cost for each neighbour needs an update 
		for _, neighbour := range links[lowestValued] {
			valve := valves[neighbour]
			//if the cost to reach its neighbour is more than the cost to reach itself + 1
			if dijkstraTable[valve] > dijkstraTable[lowestValued] + 1 {
				//I update the cost
				dijkstraTable[valve] = dijkstraTable[lowestValued] + 1 
			}
		}
		//Remove from unvisited list
		unvisited = remove(unvisited, lowestValued)
	}
	return
}

func getLowestValued(dijkstraTable map[Valve]int, visited []Valve) Valve {
	var minValve Valve
	var minValue int = 9999
	for key, value := range dijkstraTable {
		if value < minValue && !contains(visited, key) {
			minValue = value
			minValve = key
		}
	}
	return minValve
} 

func remove(list []Valve, s Valve) []Valve {
	for i, item := range list {
		if item == s {
			return append(list[:i], list[i + 1:]...)
		}
	}
	return []Valve{}
}

func initTable(valves []Valve) (dijkstraTable map[Valve]int) {
	dijkstraTable = make(map[Valve]int)
	for _, valve := range valves {
		dijkstraTable[valve] = 9999
	}
	return
}

func listItems(valves map[string]Valve) (list []Valve) {
	for _, valve := range valves {
		list = append(list, valve)
	}
	return 
}

func contains(list []Valve, s Valve) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}

//Returns:
// - a map of Valves, in which the key is the name of the valve and the value is the Valve
// - a map of Links, in which the key is the name of the valve and the value is a list of the connected valves
func parseInput() (valves map[string]Valve, links map[Valve][]string) {
	valves = make(map[string]Valve)
	links = make(map[Valve][]string)
	var name string
	var rate int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		firstComma := strings.IndexRune(line, ',')
		if firstComma < 0 {
			firstComma = len(line)
		}
		destinationsList := strings.Split(line[firstComma - 2:], ",")
		firstSemiColon := strings.IndexRune(line, ';')
		fmt.Sscanf(line[: firstSemiColon], "Valve %s has flow rate=%d", &name, &rate)
		valve := Valve{name, rate}
		valves[name] = valve
		for _, neighbour := range destinationsList {
			links[valve] = append(links[valve], strings.TrimSpace(neighbour))
		}
	}
	return valves, links
}