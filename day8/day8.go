package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var N_CONTAINS = []string{
	"abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"}

type digit struct {
	pts    string
	number int
}

type digitIo struct {
	input  []digit
	output []digit
}

func newDigit(pts string) digit {
	return digit{pts: pts, number: getNum(pts)}
}

func getNum(pts string) int {
	switch len(pts) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	default:
		return -1
	}
}

func getData(fp string) []*digitIo {
	f, _ := os.Open(fp)
	scanner := bufio.NewScanner(f)

	retIo := []*digitIo{}
	for scanner.Scan() {
		ioStr := strings.Split(scanner.Text(), " | ")
		inputStrs := strings.Fields(ioStr[0])
		outputStrs := strings.Fields(ioStr[1])

		retIo = append(retIo, &digitIo{make([]digit, 10), make([]digit, 4)})
		for i := range inputStrs {
			retIo[len(retIo)-1].input[i] = newDigit(inputStrs[i])
		}
		for i := range outputStrs {
			retIo[len(retIo)-1].output[i] = newDigit(outputStrs[i])
		}
	}
	return retIo
}

func find(input []digit, number int) int {
	for i, digit := range input {
		if digit.number == number {
			return i
		}
	}
	return -1
}

func removeFromString(a string, b string) string {
	for _, char := range b {
		a = strings.ReplaceAll(a, string(char), "")
	}
	return a
}

func removeFromInput(input []digit, number int) []digit {
	num := find(input, number)
	if num >= 0 {
		if find(input, 8) < len(input)-1 {
			return append(input[:num], input[num+1:]...)
		} else {
			return append(input[:num])
		}
	}
	return input
}

func contains(a string, b string, rearrangement bool) bool {
	for _, char := range b {
		contains := false
		for _, oChar := range a {
			if char == oChar {
				contains = true
				a = removeFromString(a, string(oChar))
				break
			}
		}
		if !contains {
			return false
		}
	}
	return !rearrangement || a == ""
}

func isRearrangement(a string, b string) bool {
	return contains(a, b, true)
}

func setActivations(input []digit, digitFunc func(digit) bool) {
	for _, digit := range input {
		if digitFunc(digit) {
			return
		}
	}
}

func findActivationMapping(input []digit) map[byte]byte {
	activations := make(map[byte]byte)
	possibilities := make(map[string][]byte)
	input_t := make([]digit, len(input))
	copy(input_t, input)

	// 1. Get one and seven, find a, and cf candidates.
	one, seven := input_t[find(input_t, 1)], input_t[find(input_t, 7)]
	activations['a'] = removeFromString(seven.pts, one.pts)[0]
	possibilities["cf"] = []byte{one.pts[0], one.pts[1]}

	// 2. Get bd candidates from one and four.
	four := input_t[find(input_t, 4)]
	dbStr := removeFromString(four.pts, one.pts)
	possibilities["db"] = []byte{dbStr[0], dbStr[1]}

	// 2.5. Remove 8, 1, 7, and 4 from input
	input_t = removeFromInput(input_t, 8)
	input_t = removeFromInput(input_t, 1)
	input_t = removeFromInput(input_t, 7)
	input_t = removeFromInput(input_t, 4)

	// 3. Find 5 as of length 5 and containing a, b, c or f, d, g.
	// Get c, f, g.
	remArr := [][]byte{append([]byte{activations['a'], possibilities["cf"][0]}, possibilities["db"]...),
		append([]byte{activations['a'], possibilities["cf"][1]}, possibilities["db"]...)}
	setActivations(input_t, (func(digit digit) bool {
		for i := range remArr {
			if (len(digit.pts) == 5) && contains(digit.pts, string(remArr[i][:]), false) {
				activations['g'] = removeFromString(digit.pts, string(remArr[i][:]))[0]
				activations['f'] = remArr[i%2][1]
				activations['c'] = remArr[(i+1)%2][1]
				return true
			}
		}
		return false
	}))

	// 4. Find 3 as of length 5 and containing a, c, d, f, g.
	remArr = [][]byte{{possibilities["db"][0], activations['a'], activations['c'], activations['f'], activations['g']},
		{possibilities["db"][1], activations['a'], activations['c'], activations['f'], activations['g']}}
	setActivations(input_t, (func(digit digit) bool {
		for i := range remArr {
			if len(digit.pts) == 5 && contains(digit.pts, string(remArr[i][:]), false) {
				activations['d'] = remArr[i%2][0]
				activations['b'] = remArr[(i+1)%2][0]
				return true
			}
		}
		return false
	}))

	// 5. Get last remaining digit
	actArr := []byte{}
	for key := range activations {
		actArr = append(actArr, activations[key])
	}
	activations['e'] = removeFromString("abcdefg", string(actArr[:]))[0]

	return reverseBijectiveMapping(activations)
}

func reverseBijectiveMapping(reverseMap map[byte]byte) map[byte]byte {
	retMap := make(map[byte]byte)
	for key, val := range reverseMap {
		retMap[val] = key
	}
	return retMap
}

func getNumMapped(num digit, activations map[byte]byte) int {
	actPts := ""
	for i := range num.pts {
		actPts += string(activations[num.pts[i]])
	}
	for i := range N_CONTAINS {
		if isRearrangement(actPts, N_CONTAINS[i]) {
			return i
		}
	}
	return -1
}

func getOutputVal(board digitIo) int {
	activations := findActivationMapping(board.input)
	actOuts := 0
	for i := range board.output {
		actOuts = actOuts*10 + getNumMapped(board.output[i], activations)
	}
	return actOuts
}

func main() {
	boardIoData := getData("part2.data")

	uniqueNumsSum := 0
	for _, ioRow := range boardIoData {
		for _, digit := range ioRow.output {
			if digit.number != -1 {
				uniqueNumsSum++
			}
		}
	}
	fmt.Printf("Number of 1, 4, 7, or 8 digits in output values: %d\n", uniqueNumsSum)
	outSum := 0
	for _, row := range boardIoData {
		outSum += getOutputVal(*row)
	}
	fmt.Printf("Sum of all outputs: %d\n", outSum)
}
