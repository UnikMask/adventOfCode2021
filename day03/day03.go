package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("part1.data")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	gammaRate := ""
	epsilonRate := ""

	// Find gamma rate and epsilon rate.
	var n [][]int
	nums := []string{}
	for scanner.Scan() {
		s := scanner.Text()
		nums = append(nums, scanner.Text())
		if n == nil {
			n = make([][]int, len(s))
		}
		for i, bit := range s {
			if n[i] == nil {
				n[i] = make([]int, 2)
			}
			switch bit {
			case '1':
				n[i][1]++
				break
			case '0':
				n[i][0]++
				break
			}
		}
	}

	// Find most common and least common string.
	for i := 0; i < len(n); i++ {
		if n[i][1] >= n[i][0] {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	// Find oxygen rating and CO2 rating from most common and least common strings.
	contenders := [][]string{make([]string, len(nums)), make([]string, len(nums))}
	copy(contenders[0], nums)
	copy(contenders[1], nums)

	for k := 0; k < len(nums[0]); k++ {
		for i := 0; i < len(contenders); i++ {
			if len(contenders[i]) <= 1 {
				continue
			}

			n := []int{0, 0}
			for j := 0; j < len(contenders[i]); j++ {
				bit := contenders[i][j][k]
				switch bit {
				case '0':
					n[0]++
					break
				case '1':
					n[1]++
				}
			}

			mostCommon := '1'
			if n[0] > n[1] {
				mostCommon = '0'
			}

			for j := 0; j < len(contenders[i]); j++ {
				if (i == 0 && contenders[i][j][k] != byte(mostCommon)) || (i == 1 && contenders[i][j][k] == byte(mostCommon)) {
					contenders[i] = append(contenders[i][:j], contenders[i][j+1:]...)
					j--
				}
			}
		}
	}

	O2, _ := strconv.ParseUint(contenders[0][0], 2, len(contenders[0][0]))
	Co2, _ := strconv.ParseUint(contenders[1][0], 2, len(contenders[1][0]))
	gamma, _ := strconv.ParseUint(gammaRate, 2, len(gammaRate))
	epsilon, _ := strconv.ParseUint(epsilonRate, 2, len(epsilonRate))
	fmt.Printf("Power Consumption: %d\n Oxygen generator rating: %d\n ", gamma*epsilon, O2*Co2)
	fmt.Printf("Oxygen Rating: %d, CO2 rating: %d\n", O2, Co2)
}
