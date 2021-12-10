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
var completionMap = map[rune]int{'}': 3, ')': 1, ']': 2, '>': 4}

type ListNode struct {
	char rune
	next *ListNode
}

type stack struct {
	first *ListNode
}

func (s *stack) push(c rune) {
	node := &ListNode{char: c}
	if s.first != nil {
		node.next = s.first
	}
	s.first = node
}

func (s *stack) pop() rune {
	if s.first != nil {
		popped := s.first
		s.first = s.first.next
		return popped.char
	} else {
		return '\n'
	}
}

func main() {
	f, _ := os.Open("part1.data")
	scanner := bufio.NewScanner(f)
	syntaxScore := 0
	completions := []int{}
	for scanner.Scan() {
		errorStack := stack{}
		errored := true
		for _, char := range scanner.Text() {
			errored = true
			possibilities := append(chunkOpens, errorStack.pop())
			for _, i := range possibilities {
				if char == i {
					if i != possibilities[len(possibilities)-1] {
						errorStack.push(possibilities[len(possibilities)-1])
						errorStack.push(chunkCloses[i])
					}
					errored = false
					break
				}
			}
			if errored {
				syntaxScore += scoreMap[char]
				break
			}
		}
		if !errored {
			lineCompletion := 0
			for next := errorStack.pop(); next != '\n'; next = errorStack.pop() {
				lineCompletion = lineCompletion*5 + completionMap[next]
			}
			completions = append(completions, lineCompletion)
		}
	}
	fmt.Printf("Syntax Error Score: %d\n", syntaxScore)

	sort.Ints(completions)
	fmt.Printf("Completion Score: %d\n", completions[len(completions)/2])
}
