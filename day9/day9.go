package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	BOLD  = "\033[1m"
	RESET = "\033[0m"
)

type tuple struct {
	x int
	y int
}

type tupleSet map[tuple]struct{}

func (s tupleSet) add(t tuple) {
	s[t] = struct{}{}
}

func (s tupleSet) addAll(tv []tuple) {
	for _, t := range tv {
		s.add(t)
	}
}

func (s tupleSet) has(t tuple) bool {
	for key := range s {
		if key == t {
			return true
		}
	}
	return false
}

var EMPTY_TUPLE = tuple{-1, -1}

func getData(fp string) [][]int {
	f, _ := os.Open(fp)
	scanner := bufio.NewScanner(f)

	heightMap := [][]int{}
	for scanner.Scan() {
		heightMap = append(heightMap, []int{})
		for _, char := range scanner.Text() {
			height, _ := strconv.Atoi(string(char))
			heightMap[len(heightMap)-1] = append(heightMap[len(heightMap)-1], height)
		}
	}
	return heightMap
}

func isLegal(position tuple, heightMap [][]int) bool {
	return position.x >= 0 && position.x < len(heightMap[0]) && position.y >= 0 && position.y < len(heightMap)
}

func getAdjacentPts(position tuple, prevPos tuple, heightMap [][]int, depth int) []tuple {
	if depth <= 0 {
		return []tuple{position}
	} else {
		adjacentPts := []tuple{{position.x + 1, position.y}, {position.x - 1, position.y},
			{position.x, position.y + 1}, {position.x, position.y - 1}}
		for pt := 0; pt < len(adjacentPts); pt++ {
			if !isLegal(adjacentPts[pt], heightMap) || prevPos == position || heightMap[adjacentPts[pt].y][adjacentPts[pt].x] == 9 {
				adjacentPts = append(adjacentPts[:pt], adjacentPts[pt+1:]...)
				pt--
			}
		}
		nextPts := []tuple{}
		for _, pt := range adjacentPts {
			nextPts = append(nextPts, getAdjacentPts(pt, position, heightMap, depth-1)...)
		}
		return append(adjacentPts, nextPts...)
	}
}

func isLowestAdjacentPt(position tuple, heightMap [][]int) bool {
	adjacentPts := getAdjacentPts(position, EMPTY_TUPLE, heightMap, 1)
	for _, pt := range adjacentPts {
		if heightMap[pt.y][pt.x] <= heightMap[position.y][position.x] && pt != position {
			return false
		}
	}
	return true
}

func getLowPtsLocations(heightMap [][]int) map[tuple]int {
	lowPtsMap := make(map[tuple]int)
	for i, row := range heightMap {
		for j, pt := range row {
			if isLowestAdjacentPt(tuple{j, i}, heightMap) {
				lowPtsMap[tuple{j, i}] = pt
			}
		}
	}
	return lowPtsMap
}

func getBasin(position tuple, heightMap [][]int) tupleSet {
	s := tupleSet{}
	if heightMap[position.y][position.x] == 9 {
		return s
	}
	s.add(position)

	nextPts := []tuple{position}
	for len(nextPts) > 0 {
		newPts := []tuple{}
		for _, pt := range nextPts {
			curPtAdjacencies := []tuple{{pt.x + 1, pt.y}, {pt.x - 1, pt.y},
				{pt.x, pt.y + 1}, {pt.x, pt.y - 1}}
			for _, adjPt := range curPtAdjacencies {
				if isLegal(adjPt, heightMap) && !s.has(adjPt) {
					if heightMap[adjPt.y][adjPt.x] != 9 {
						s.add(adjPt)
						newPts = append(newPts, adjPt)
					}
				}
			}
		}
		nextPts = newPts
	}

	return s
}

func isInBasin(position tuple, heightMap [][]int, basins []tupleSet) bool {
	if heightMap[position.y][position.x] == 9 {
		return false
	}
	for _, basin := range basins {
		for pt := range basin {
			if position == pt {
				return true
			}
		}
	}
	return false
}

func main() {
	heightMap := getData("part1.data")
	lowPtsMap := getLowPtsLocations(heightMap)

	riskSum := 0
	for _, height := range lowPtsMap {
		riskSum += height + 1
	}

	basins := []tupleSet{}
	for i := range heightMap {
		prevPosTop := true
		for j := range heightMap[i] {
			if heightMap[i][j] == 9 {
				prevPosTop = true
			} else if prevPosTop {
				if !isInBasin(tuple{j, i}, heightMap, basins) {
					basins = append(basins, getBasin(tuple{j, i}, heightMap))
					prevPosTop = false
				}
			}
		}
	}
	fmt.Printf("Sum of risk levels of all low points: %d\n", riskSum)

	topBasins := []tupleSet{{}, {}, {}}
	for _, basin := range basins {
		for i := range topBasins {
			if len(basin) > len(topBasins[i]) {
				for j := len(topBasins) - 1; j > i; j-- {
					topBasins[j] = topBasins[j-1]
				}
				topBasins[i] = basin
				break
			}
		}
	}
	fmt.Printf("Product of top 3 basins: %d\n", len(topBasins[0])*len(topBasins[1])*len(topBasins[2]))
}
