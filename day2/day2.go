package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	dist  int
	depth int
	next  *pos
}

func getInstruction(s string) (*pos, error) {
	instruction := strings.Split(s, " ")
	spd, err := strconv.Atoi(instruction[1])
	if err != nil {
		return nil, err
	} else {
		retPos := &pos{}
		switch instruction[0] {
		case "forward":
			retPos.dist = spd
			break
		case "up":
			retPos.depth = -spd
			break
		case "down":
			retPos.depth = spd
			break
		default:
			return nil, errors.New("Invalid instruction format!")
		}
		return retPos, nil
	}
}

func getData(filename string) (*pos, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Get 1st line
	scanner.Scan()
	retPos, err := getInstruction(scanner.Text())
	if err != nil {
		return nil, err
	}
	posPtr := retPos

	// Get next lines
	for scanner.Scan() {
		posPtr.next, err = getInstruction(scanner.Text())
		if err != nil {
			return nil, err
		} else {
			posPtr = posPtr.next
		}
	}
	return retPos, nil
}

func getPos2(posList *pos, aim int) (int, int) {
	retDist, retDepth := 0, 0
	if posList.next != nil {
		retDist, retDepth = getPos2(posList.next, aim+posList.depth)
	}
	retDist += posList.dist
	retDepth += aim * posList.dist
	return retDist, retDepth
}

func getSums(posList *pos) (int, int) {
	retDist, retDepth := 0, 0
	if posList.next != nil {
		retDist, retDepth = getSums(posList.next)
	}
	retDist += posList.dist
	retDepth += posList.depth
	return retDist, retDepth
}

func main() {
	posList, err := getData("part1.data")
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	dist, depth := getSums(posList)
	fmt.Printf("-- Part 1 --\n Dist %d, depth %d, mult %d\n", dist, depth, dist*depth)

	dist2, depth2 := getPos2(posList, 0)
	fmt.Printf("-- Part 2 --\nDist %d, depth %d, mult %d\n", dist2, depth2, dist2*depth2)
}
