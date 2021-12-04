package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	STATE_WIN     = 1
	STATE_NEUTRAL = 0
	STATE_ERR     = -1
)

type Board struct {
	board   [][]int
	checked []tuple
	dims    tuple
	done    bool
}

type tuple struct {
	x int
	y int
}

func (self *Board) At(pos tuple) int {
	if !(pos.x >= self.dims.x) && !(pos.x < 0) && !(pos.y >= self.dims.y) && !(pos.y < 0) {
		return self.board[pos.y][pos.x]
	} else {
		return -1
	}
}

func newBoard(board [][]int, dims *tuple) *Board {
	retBoard := &Board{board: board, dims: *dims}
	retBoard.checked = []tuple{}
	retBoard.done = false
	return retBoard
}

func (self *Board) isChecked(pos *tuple) bool {
	for i := 0; i < len(self.checked); i++ {
		if self.checked[i].x == pos.x && self.checked[i].y == pos.y {
			return true
		}
	}
	return false
}

func (self *Board) checkState() int {
	if self.done {
		return STATE_WIN
	}

	// Check rows
	check := false
	for i := 0; i < self.dims.y; i++ {
		check = true
		for j := 0; j < self.dims.x; j++ {
			if !self.isChecked(&tuple{x: j, y: i}) {
				check = false
				break
			}
		}
		if check {
			break
		}
	}
	if !check {
		for i := 0; i < self.dims.x; i++ {
			check = true
			for j := 0; j < self.dims.y; j++ {
				if !self.isChecked(&tuple{x: i, y: j}) {
					check = false
					break
				}
			}
			if check {
				break
			}
		}
	}
	if check {
		self.done = true
		return STATE_WIN
	} else {
		return STATE_NEUTRAL
	}
}

func (self *Board) checkNb(n int) bool {
	for i := 0; i < self.dims.y; i++ {
		for j := 0; j < self.dims.x; j++ {
			pos := &tuple{x: j, y: i}
			if self.At(*pos) == n {
				self.checked = append(self.checked, *pos)
				return true
			}
		}
	}
	return false
}

func (self *Board) getScore(calledNum int) int {
	origScore := 0
	for y, row := range self.board {
		for x, n := range row {
			if !self.isChecked(&tuple{x: x, y: y}) {
				origScore += n
			}
		}
	}
	return origScore * calledNum
}

func getData(filepath string) ([]int, []*Board) {
	f, _ := os.Open(filepath)
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	rndNumsStr := strings.Split(scanner.Text(), ",")

	// Get random numbers
	rndNums := make([]int, len(rndNumsStr))
	for i, n_str := range rndNumsStr {
		n, _ := strconv.Atoi(n_str)
		rndNums[i] = n
	}

	// Get boards
	boardArr := []*Board{}
	var curBoard [][]int
	for scanner.Scan() {
		if scanner.Text() == "" {
			if curBoard != nil {
				dims := &tuple{x: len(curBoard[0]), y: len(curBoard)}
				boardArr = append(boardArr, newBoard(curBoard, dims))
			}
			curBoard = [][]int{}
		} else {
			rowStr := strings.Fields(scanner.Text())
			row := make([]int, len(rowStr))
			for i, n_str := range rowStr {
				n, _ := strconv.Atoi(n_str)
				row[i] = n
			}
			curBoard = append(curBoard, row)
		}
	}
	return rndNums, boardArr
}

func main() {
	rndNums, boardArr := getData("part1.data")

	first := true
	var lastBoardWinScore int
	for i := 0; i < len(rndNums); i++ {
		calledNum := rndNums[i]
		winningBoards := make(map[int]*Board)

		// Set numbers for all boards.
		for j, board := range boardArr {
			if board.checkNb(calledNum) && !board.done {
				winningBoards[j] = board
			}
		}

		// Check condition for all winning boards
		for j, board := range winningBoards {
			if board.checkState() == STATE_WIN {
				lastBoardWinScore = board.getScore(calledNum)
				if first {
					fmt.Printf("Board no. %d - Winning board score: %d\n", j, lastBoardWinScore)
					first = false
				}
			}
		}
	}
	fmt.Printf("Last Board Winning score: %d\n", lastBoardWinScore)
}
