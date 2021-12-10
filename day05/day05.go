package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tuple struct {
	x int
	y int
}

func (a tuple) times(n int) tuple {
	return tuple{x: a.x * n, y: a.y * n}
}

func (a tuple) plus(b tuple) tuple {
	return tuple{x: a.x + b.x, y: a.y + b.y}
}

func (a tuple) equals(b tuple) bool {
	return a.x == b.x && a.y == b.y
}

func appendToLineMap(lineMap map[tuple]int, t tuple) map[tuple]int {
	_, exists := lineMap[t]
	if !exists {
		lineMap[t] = 1
	} else {
		lineMap[t] += 1
	}
	return lineMap
}

func getMap(filepath string, diagonals bool) map[tuple]int {
	f, _ := os.Open(filepath)

	scanner := bufio.NewScanner(f)
	lineMap := make(map[tuple]int)

	// Scan file
	for scanner.Scan() {
		// Get endpoints for lines
		endpointStrs := strings.Split(scanner.Text(), "->")
		endpoints := [2]tuple{{0, 0}, {0, 0}}
		for i := range endpoints {
			endpointSingleStr := strings.Split(endpointStrs[i], ",")
			endpoints[i].x, _ = strconv.Atoi(strings.ReplaceAll(endpointSingleStr[0], " ", ""))
			endpoints[i].y, _ = strconv.Atoi(strings.ReplaceAll(endpointSingleStr[1], " ", ""))
		}

		// Set line given horizontal or vertical line
		step := tuple{1, 1}
		if diagonals || endpoints[0].x == endpoints[1].x || endpoints[0].y == endpoints[1].y {
			if endpoints[0].x == endpoints[1].x {
				step.x = 0
			} else if endpoints[0].x > endpoints[1].x {
				step.x = -1
			} else {
				step.x = 1
			}

			if endpoints[0].y == endpoints[1].y {
				step.y = 0
			} else if endpoints[0].y > endpoints[1].y {
				step.y = -1
			} else {
				step.y = 1
			}

			// Draw line

			lineMap = appendToLineMap(lineMap, endpoints[1])
			for i := 0; !endpoints[0].plus(step.times(i)).equals(endpoints[1]); i++ {
				lineMap = appendToLineMap(lineMap, endpoints[0].plus(step.times(i)))
			}
		}
	}
	return lineMap
}

func getNbCollidingSpots(lineMap map[tuple]int, limit int) int {
	n := 0
	for _, val := range lineMap {
		if val >= limit {
			n++
		}
	}
	return n
}

func main() {
	fmt.Printf("Number of 2-overlaps in no-diagonal diagram: %d\n ", getNbCollidingSpots(getMap("part1.data", false), 2))
	fmt.Printf("Number of 2-overlaps in diagram: %d\n ", getNbCollidingSpots(getMap("part1.data", true), 2))
}
