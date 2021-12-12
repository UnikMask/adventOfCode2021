package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var chunkOpens = []rune{'{', '[', '(', '<'}
var chunkCloses = map[rune]rune{'{': '}', '[': ']', '(': ')', '<': '>'}
var scoreMap = map[rune]int{'}': 1197, ')': 3, ']': 57, '>': 25137}
var completionMap = map[rune]int{'}': 3, ')': 1, ']': 2, '>': 4, '\n': 0}

func main() {
	f, _ := os.Open("part1.data")
	scanner := bufio.NewScanner(f)
	syntaxScore := 0
	completions := []int{}
file:
	for scanner.Scan() {
		errorStack := []rune{0}
	line:
		for _, char := range scanner.Text() {
			for _, i := range chunkOpens {
				if char == errorStack[0] {
					errorStack = errorStack[1:]
					continue line
				} else if i == char {
					errorStack = append([]rune{chunkCloses[i]}, errorStack...)
					continue line
				}
			}
			syntaxScore += scoreMap[char]
			continue file
		}
		lineCompletion := 0
		for _, next := range errorStack {
			lineCompletion = lineCompletion*5 + completionMap[next]
		}
		completions = append(completions, lineCompletion)
	}
	sort.Ints(completions)
	fmt.Printf("Syntax Error Score: %d\n", syntaxScore)
	fmt.Printf("Completion Score: %d\n", completions[len(completions)/2])
}
