package main
import ("bufio"; "fmt"; "os"; "sort")
func main() {
	f, _ := os.Open("part1.data")
	sc, score, compl, closes := bufio.NewScanner(f), 0, []int{}, map[rune]rune{'{': '}', '[': ']', '(': ')', '<': '>'}
	scoreMap := map[rune]struct{error int; completion int}{'}': {1197, 3}, ')': {3, 1}, ']': {57, 2}, '>': {25137, 4}}
file: for sc.Scan() {
		errorStack := []rune{'\n'}
		for _, char := range sc.Text() {
			if char == errorStack[0] {
				errorStack = errorStack[1:]
			} else if v, exists := closes[char]; exists {
				errorStack = append([]rune{v}, errorStack...)
			} else {
				score += scoreMap[char].error
				continue file
			}
		}
		compl = append(compl, 0)
		for i := 0; i < len(errorStack)-1; i++ {
			compl[len(compl)-1] = compl[len(compl)-1]*5 + scoreMap[errorStack[i]].completion
		}
	}
	sort.Ints(compl)
	fmt.Printf("Syntax Error Score: %d\nCompletion Score: %d\n", score, compl[len(compl)/2])
}
