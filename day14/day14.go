package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type tuple struct {
	c rune
	v int
}

type tupleList []tuple

func (p tupleList) Len() int           { return len(p) }
func (p tupleList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p tupleList) Less(i, j int) bool { return p[i].v > p[j].v }

func nextStep(pairMap map[string]int, commonsMap map[rune]int, insertions map[string]string) (map[string]int, map[rune]int) {
	retPairMap := map[string]int{}
	retCommons := map[rune]int{}
	for k, v := range commonsMap {
		retCommons[k] = v
	}
	for pair, v := range pairMap {
		retPairMap[pair] += v
		if nc, ok := insertions[pair]; ok {
			retPairMap[string(pair[0])+nc] += v
			retPairMap[nc+string(pair[1])] += v
			retCommons[rune(nc[0])] += v
			retPairMap[pair] -= v
		}
	}
	return retPairMap, retCommons
}

func main() {
	f, _ := os.Open("part1.data")
	sc := bufio.NewScanner(f)
	sc.Scan()
	template := sc.Text()
	insertions := map[string]string{}
	for sc.Scan() {
		if sc.Text() != "" {
			ends := strings.Split(sc.Text(), " -> ")
			insertions[ends[0]] = ends[1]
		}
	}
	pairMap := map[string]int{}
	commonsMap := map[rune]int{}
	for _, c := range template {
		commonsMap[c]++
	}
	for i := 1; i < len(template); i++ {
		pairMap[template[i-1:i+1]]++
	}
	steps := []struct {
		pm map[string]int
		cm map[rune]int
	}{{pairMap, commonsMap}}
	for i := 0; i < 40; i++ {
		steps = append(steps, struct {
			pm map[string]int
			cm map[rune]int
		}{map[string]int{}, map[rune]int{}})
		steps[i+1].pm, steps[i+1].cm = nextStep(steps[i].pm, steps[i].cm, insertions)
	}

	cl10, cl40 := tupleList{}, tupleList{}
	for k, v := range steps[10].cm {
		cl10 = append(cl10, tuple{k, v})
	}
	for k, v := range steps[40].cm {
		cl40 = append(cl40, tuple{k, v})
	}
	sort.Sort(cl10)
	sort.Sort(cl40)
	fmt.Printf("Part 1: %d\nPart 2: %d\n", cl10[0].v-cl10[len(cl10)-1].v, cl40[0].v-cl40[len(cl40)-1].v)
}
