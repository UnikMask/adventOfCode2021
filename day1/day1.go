package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type depth struct {
	depth int
	next  *depth
}

func getData(file string) (*depth, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	retDepth := &depth{}
	scanner := bufio.NewScanner(f)

	// Get 1st line.
	scanner.Scan()
	baseDepth, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}
	retDepth = &depth{depth: baseDepth}
	var depthPtr *depth = retDepth

	// Scan for each line.
	for scanner.Scan() {
		nextDepth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		depthPtr.next = &depth{depth: nextDepth}
		depthPtr = depthPtr.next
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return retDepth, nil
}

func (d *depth) getSumOf(sumOf int) (int, bool) {
	return d.getSumOfRec(0, sumOf)
}

func (d *depth) getSumOfRec(current int, sumOf int) (int, bool) {
	if sumOf <= 0 {
		return d.depth + current, true
	} else if d.next == nil {
		return 0, false
	} else {
		return d.next.getSumOfRec(current+d.depth, sumOf-1)
	}
}

func (d *depth) getNumIncreases(sumOf int) int {
	return d.getNumIncreasesRec(sumOf, 0)
}

func (d *depth) getNumIncreasesRec(sumOf int, increases int) int {
	sum, exists := d.getSumOf(sumOf)
	if !exists || d.next == nil {
		return increases
	}

	nextSum, exists := d.next.getSumOf(sumOf)
	if !exists {
		return increases
	} else if nextSum-sum > 0 {
		increases++
	}
	return d.next.getNumIncreasesRec(sumOf, increases)

}

func main() {
	depthMap, err := getData("part1.data")
	if err != nil {
		log.Fatal(err)
	}

	// Read depth list to find increases
	log.Printf("Number of depth increases: %d", depthMap.getNumIncreases(0))
	log.Printf("Number of depth increases at sum 3: %d", depthMap.getNumIncreases(2))
}
