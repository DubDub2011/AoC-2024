package main

import (
	"bufio"
	"day2/solution"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()

	fmt.Printf("Part 1: %d\n", solution.ReportSafeCount(input))
	fmt.Printf("Part 2: %d\n", solution.ReportSafeCount2(input))
}

func readInput() [][]int {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	result := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		var lineRes []int
		for _, part := range parts {
			val, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			lineRes = append(lineRes, val)
		}

		result = append(result, lineRes)
	}

	return result
}
