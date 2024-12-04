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
		safeLines += isSafe(levels)
	}
	fmt.Println(safeLines)
}


func isSafe(levels []int) int {
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