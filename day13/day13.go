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

func fold(grid map[tuple]struct{}, instrc tuple) map[tuple]struct{} {
	ret := map[tuple]struct{}{}
	for k, v := range grid {
		ret[k] = v
	}
	deleteSet := map[tuple]struct{}{}
	if instrc.x != 0 {
		for k := range grid {
			if k.x > instrc.x {
				ret[tuple{2*instrc.x - k.x, k.y}] = struct{}{}
				deleteSet[k] = struct{}{}
			}
		}
		goto delete
	}
	for k := range grid {
		if k.y > instrc.y {
			ret[tuple{k.x, 2*instrc.y - k.y}] = struct{}{}
			deleteSet[k] = struct{}{}
		}
	}
delete:
	for k := range deleteSet {
		delete(ret, k)
	}
	return ret
}

func main() {
	f, _ := os.Open("part1.data")
	sc := bufio.NewScanner(f)
	grid := map[tuple]struct{}{}
	instructions := []tuple{}

	getInstructions := func(instrcs []tuple, grid map[tuple]struct{}, text string) ([]tuple, map[tuple]struct{}) {
		foldStr := strings.Split(strings.ReplaceAll(text, "fold along ", ""), "=")
		foldn, _ := strconv.Atoi(foldStr[1])
		switch foldStr[0] {
		case "x":
			return append(instrcs, tuple{foldn, 0}), grid
		default:
			return append(instrcs, tuple{0, foldn}), grid
		}
	}

	getGridPos := func(instrcs []tuple, grid map[tuple]struct{}, text string) ([]tuple, map[tuple]struct{}) {
		x, _ := strconv.Atoi(strings.Split(sc.Text(), ",")[0])
		y, _ := strconv.Atoi(strings.Split(sc.Text(), ",")[1])
		grid[tuple{x, y}] = struct{}{}
		return instrcs, grid

	}

	dataFunc := getGridPos
	for sc.Scan() {
		if sc.Text() != "" {
			instructions, grid = dataFunc(instructions, grid, sc.Text())
			continue
		}
		dataFunc = getInstructions
	}
	fmt.Printf("Number of visible dots at no fold: %d\n", len(grid))
	gridBase := fold(grid, instructions[0])
	for _, instrc := range instructions {
		grid = fold(grid, instrc)
	}
	minMax := tuple{0, 0}
	for k := range grid {
		if minMax.x < k.x {
			minMax.x = k.x
		}
		if minMax.y < k.y {
			minMax.y = k.y
		}
	}
	fmt.Printf("Number of visible dots: %d\nNumber of visible dots after all folds: %d\n", len(gridBase), len(grid))
	for i := 0; i <= minMax.y; i++ {
		for j := 0; j <= minMax.x; j++ {
			if _, ok := grid[tuple{j, i}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}

}
