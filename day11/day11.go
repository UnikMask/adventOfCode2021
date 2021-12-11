package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type tuple struct {
	x int
	y int
}

func isLegal(p tuple, octopuses [][]int) bool {
	return p.x >= 0 && p.x < len(octopuses[0]) && p.y >= 0 && p.y < len(octopuses)
}

func addAll(a map[tuple]struct{}, b map[tuple]struct{}) map[tuple]struct{} {
	ret := make(map[tuple]struct{})
	for key, val := range a {
		ret[key] = val
	}
	for key, val := range b {
		ret[key] = val
	}
	return ret
}

func combine(a map[tuple]int, b map[tuple]int) map[tuple]int {
	ret := make(map[tuple]int)
	for k, v := range a {
		ret[k] = v
	}
	for k, v := range b {
		if _, exists := ret[k]; exists {
			ret[k] += v
		} else {
			ret[k] = v
		}
	}
	return ret
}

func getFriends(position tuple, octopuses [][]int) map[tuple]int {
	retSet := make(map[tuple]int)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if isLegal(tuple{x: position.x + j, y: position.y + i}, octopuses) {
				retSet[tuple{x: position.x + j, y: position.y + i}] = 1
			}
		}
	}
	return retSet
}

func flash(flashSet map[tuple]struct{}, octopuses [][]int, exploded map[tuple]struct{}, friendsSet map[tuple]int) (map[tuple]struct{}, map[tuple]int) {
	exploded = addAll(exploded, flashSet)
	for p := range flashSet {
		friendsSet = combine(friendsSet, getFriends(p, octopuses))
	}
	for p := range exploded {
		if _, exists := friendsSet[p]; exists {
			delete(friendsSet, p)
		}
	}
	flashSet = make(map[tuple]struct{})
	friendDelete := []tuple{}
	for p, add := range friendsSet {
		if octopuses[p.y][p.x]+add > 9 {
			flashSet[p] = struct{}{}
			friendDelete = append(friendDelete, p)
		}
	}
	for _, p := range friendDelete {
		delete(friendsSet, p)
	}
	if len(flashSet) > 0 {
		return flash(flashSet, octopuses, exploded, friendsSet)
	}
	return exploded, friendsSet
}

func nextStep(octopuses [][]int) ([][]int, int) {
	ret := make([][]int, len(octopuses))

	// Increase value of all octopuses by 1, get flashing octopuses.
	flashSet := make(map[tuple]struct{})
	for i, row := range octopuses {
		ret[i] = make([]int, len(octopuses[i]))
		for j, octopus := range row {
			ret[i][j] = octopus + 1
			if ret[i][j] > 9 {
				flashSet[tuple{x: j, y: i}] = struct{}{}
			}
		}
	}
	exploded, neighbours := flash(flashSet, ret, map[tuple]struct{}{}, map[tuple]int{})
	for p := range exploded {
		ret[p.y][p.x] = 0
	}
	for p, v := range neighbours {
		ret[p.y][p.x] += v
	}
	return ret, len(exploded)
}

func main() {
	f, _ := os.Open("part1.data")
	scanner := bufio.NewScanner(f)
	octopuses := [][]int{}
	for i := 0; scanner.Scan(); i++ {
		octopuses = append(octopuses, []int{})
		for _, octopus := range scanner.Text() {
			a, _ := strconv.Atoi(string(octopus))
			octopuses[i] = append(octopuses[i], a)
		}
	}

	totalFlashes := 0
	octopusesSync := octopuses
	for i := 0; i < 100; i++ {
		octopusesStep, incr := nextStep(octopuses)
		octopuses = octopusesStep
		totalFlashes += incr
	}
	fmt.Printf("Total number of flashes: %d\n", totalFlashes)

	stepNum := 1
	for step, inc := nextStep(octopusesSync); inc < len(octopuses)*len(octopuses[0]); step, inc = nextStep(step) {
		stepNum++
	}
	fmt.Printf("First step with full synchronisation: %d\n", stepNum)

}
