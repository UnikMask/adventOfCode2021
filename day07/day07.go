package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getData(fp string) []int {
	f, _ := os.Open(fp)
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	crabsStr := strings.Split(scanner.Text(), ",")
	crabs := make([]int, len(crabsStr))
	for i := range crabsStr {
		crabs[i], _ = strconv.Atoi(crabsStr[i])
	}
	return crabs
}

func fuelCost(crabs []int, target int) int {
	sum := 0
	for i := range crabs {
		sum += int(math.Abs(float64(crabs[i] - target)))
	}
	return sum
}

func distFuelCost(crabs []int, target int) int {
	sum := 0
	for i := range crabs {
		dist := int(math.Abs(float64(crabs[i] - target)))
		for j := 1; j <= dist; j++ {
			sum += j
		}
	}
	return sum
}

func getMedian(arr []int) int {
	arrSlice := arr[:]
	sort.Ints(arrSlice)
	return arrSlice[len(arrSlice)/2]
}

func getAvg(arr []int) (int, int) {
	sum := 0.0
	for i := range arr {
		sum += float64(arr[i])
	}
	sum /= float64(len(arr))
	return int(sum), int(sum + 0.5)
}

func getMinDistFuelCost(crabs []int) int {
	minAvg, maxAvg := getAvg(crabs)
	return int(math.Min(float64(distFuelCost(crabs, minAvg)), float64(distFuelCost(crabs, maxAvg))))
}

func main() {
	crabs := getData("part1.data")
	fmt.Printf("Lowest fuel consumption: %d\n", fuelCost(crabs, getMedian(crabs)))
	fmt.Printf("Lowest fuel consumption w/ dist augmentation: %d\n", getMinDistFuelCost(crabs))
}
