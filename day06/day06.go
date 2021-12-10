package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func nextDay(creatures []int) []int {
	retCreatures := make([]int, len(creatures))
	for i := range creatures {
		if i == 0 {
			retCreatures[8] = creatures[i]
		}
		if i < 7 {
			retCreatures[(i+6)%7] += creatures[i]
		} else {
			retCreatures[i-1] += creatures[i]
		}
	}
	return retCreatures
}

func sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	} else if len(arr) == 1 {
		return arr[0]
	} else {
		return arr[0] + sum(arr[1:])
	}
}

func getData(filepath string) []int {
	f, _ := os.Open(filepath)
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	numsStr := strings.Split(scanner.Text(), ",")
	creatures := make([]int, 9)
	for i := range numsStr {
		num, _ := strconv.Atoi(numsStr[i])
		creatures[num]++
	}
	return creatures
}

func main() {
	creatureState := make([][]int, 257)
	creatureState[0] = getData("part1.data")
	for i := 1; i < len(creatureState); i++ {
		creatureState[i] = nextDay(creatureState[i-1])
	}
	fmt.Printf("Number of creatures after 80 days: %d\n", sum(creatureState[80]))
	fmt.Printf("Number of creatures after 256 days: %d\n", sum(creatureState[256]))

}
