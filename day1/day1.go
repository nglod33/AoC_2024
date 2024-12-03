package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	firstLine := []int{}
	secondLine := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		val1, _ := strconv.Atoi(numbers[0])
		val2, _ := strconv.Atoi(numbers[1])
		firstLine = append(firstLine, val1)
		secondLine = append(secondLine, val2)
	}

	firstLine = sort(firstLine)
	secondLine = sort(secondLine)

	diff := 0

	for i := 0; i < len(firstLine); i++ {
		diff += int(math.Abs(float64(firstLine[i] - secondLine[i])))
	}

	fmt.Println(diff)

}

func sort(input []int) []int {
	if len(input) <= 1 {
		return input
	} else {
		middle := input[0]
		less := []int{}
		more := []int{}
		for i := 1; i < len(input); i++ {
			if input[i] <= middle {
				less = append(less, input[i])
			} else {
				more = append(more, input[i])
			}
		}
		return slices.Concat(sort(less), []int{middle}, sort(more))
	}
}
