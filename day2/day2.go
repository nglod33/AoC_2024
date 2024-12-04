package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"math"
	"fmt"
	"strconv"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	safeLines := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		var levels []int
		for _, part := range parts {
			level, _ := strconv.Atoi(part)
			levels = append(levels, level)
		}
		safe := isSafe(levels)
		if safe == 0 {
			// All but the first index
			safe += isSafeStrict(levels[1:])
		}
		safeLines += safe
	}
	fmt.Println(safeLines)
}

// Strict version for if the first is the problem
func isSafeStrict(levels []int) int {
	if len(levels) <= 1 {
		return 0
	}
	increasing := levels[1] > levels[0]
	for i := 1; i < len(levels); i++ {
		if increasing != (levels[i] > levels[i-1]) {
			return 0
		}
		diff := int(math.Abs(float64(levels[i-1]) - float64(levels[i])))
		if diff > 3 || diff < 1 {
			return 0
		}
	}
	return 1
}

// Doesn't work if first one is the problem
func isSafe(levels []int) int {
	err := 0
	gap := 0
	if len(levels) <= 1 {
		return 0
	}

	// str := ""
	increasing := levels[1] > levels[0]
	for i := 1; i < len(levels); i++ {
		// str = fmt.Sprintf("i: %d", i)
		// fmt.Println(str)
		// str = fmt.Sprintf("increasing: %t", increasing)
		// fmt.Println(str)
		// fmt.Println(levels[i])
		// fmt.Println(gap)
		// fmt.Println(levels[i-1-gap])
		// fmt.Println("Break")
		if increasing != (levels[i] > levels[i-1-gap]) {
			err += 1
			if i == 2 {
				increasing = levels[2] > levels[0]
				gap = 0
				continue
			}
			if err > 1 {
				return 0
			}
			gap = 1
			continue
		}
		diff := int(math.Abs(float64(levels[i-1-gap]) - float64(levels[i])))
		if diff > 3 || diff < 1 {
			err += 1
			if err > 1 {
				return 0
			}
			gap = 1
			continue
		}
		gap = 0
	}
	return 1
}