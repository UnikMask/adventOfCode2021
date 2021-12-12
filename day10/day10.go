package main
import ("bufio"; "fmt"; "os"; "sort")
var chunkCloses = map[rune]rune{'{': '}', '[': ']', '(': ')', '<': '>'}
var scoreMap = map[rune]struct{error int; completion int}{'}': {1197, 3}, ')': {3, 1}, ']': {57, 2}, '>': {25137, 4}}
func main() {
	f, _ := os.Open("part1.data")
	scanner, syntaxScore, compl := bufio.NewScanner(f), 0, []int{}
file: for scanner.Scan() {
		errorStack := []rune{'\n'}
		for _, char := range scanner.Text() {
			if char == errorStack[0] {
				errorStack = errorStack[1:]
			} else if v, exists := chunkCloses[char]; exists {
				errorStack = append([]rune{v}, errorStack...)
			} else {
				syntaxScore += scoreMap[char].error
				continue file
			}
		}
		compl = append(compl, 0)
		for i := 0; i < len(errorStack)-1; i++ {
			compl[len(compl)-1] = compl[len(compl)-1]*5 + scoreMap[errorStack[i]].completion
		}
	}
	sort.Ints(compl)
	fmt.Printf("Syntax Error Score: %d\nCompletion Score: %d\n", syntaxScore, compl[len(compl)/2])
}
