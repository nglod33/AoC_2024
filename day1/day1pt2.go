package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	scoreMap := make(map[int]int)

	for i := 0; i < len(secondLine); i++ {
		_ , exists := scoreMap[secondLine[i]]
		if exists {
			scoreMap[secondLine[i]] += 1
		} else {
			scoreMap[secondLine[i]] = 1
		}
	}

	score := 0

	for i := 0; i < len(firstLine); i++ {
		value, exists := scoreMap[firstLine[i]]
		if exists {
			score += firstLine[i] * value
		}
	}

	fmt.Println(score)

}
