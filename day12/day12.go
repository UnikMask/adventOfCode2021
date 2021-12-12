package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const (
	START = "start"
	END   = "end"
)

type caveNode struct {
	cave cave
	next []caveNode
}

type cave struct {
	name string
	big  bool
}

type connection struct {
	a cave
	b cave
}

func getCave(name string) cave {
	if !unicode.IsLower(rune(name[0])) {
		return cave{name: name, big: true}
	}
	return cave{name: name, big: false}
}

func getTree(dataMap map[connection]struct{}, lim int) caveNode {
	rootNode := caveNode{cave: cave{name: START, big: false}, next: []caveNode{}}
	rootNode.next = getTreeRec(dataMap, map[cave]int{rootNode.cave: lim}, rootNode.cave, lim)
	return rootNode
}

func getTreeRec(dataMap map[connection]struct{}, offs map[cave]int, curCave cave, lim int) []caveNode {
	next := []caveNode{}
	for con := range dataMap {
		if con.a == curCave || con.b == curCave {
			word := con.b
			if con.b == curCave {
				word = con.a
			}
			switch word.name {
			case END:
				next = append(next, caveNode{cave: word, next: nil})
				break
			default:
				if i := offs[word]; i < lim || word.big {
					newOffs := make(map[cave]int)
					for k, v := range offs {
						newOffs[k] = v
					}
					newLim := lim
					if offs[word] == lim-1 && !word.big && lim > 1 {
						newLim--
					}
					newOffs[word] += 1
					next = append(next, caveNode{cave: word, next: getTreeRec(dataMap, newOffs, word, newLim)})
				}
			}
		}
	}
	return next
}

func getPaths(root caveNode, current string) map[string]struct{} {
	arr := map[string]struct{}{}
	for _, node := range root.next {
		switch node.cave.name {
		case END:
			arr[current+"-end"] = struct{}{}
			break
		default:
			for path := range getPaths(node, current+"-"+node.cave.name) {
				arr[path] = struct{}{}
			}
		}
	}
	return arr
}

func main() {
	f, _ := os.Open("part1.data")
	scanner := bufio.NewScanner(f)

	dataMap := map[connection]struct{}{}
	for scanner.Scan() {
		ends := strings.Split(scanner.Text(), "-")
		dataMap[connection{a: getCave(ends[0]), b: getCave(ends[1])}] = struct{}{}
	}

	// Part 1
	tree := getTree(dataMap, 1)
	pathSet := getPaths(tree, "start")
	fmt.Printf("Part 1 - Number of paths: %d\n", len(pathSet))

	// Part 2
	tree = getTree(dataMap, 2)
	pathSet = getPaths(tree, "start")
	fmt.Printf("Part 2 - Number of paths: %d\n", len(pathSet))

}
