package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Input: %v", readInput())
}

func readInput() [][]int {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	res := make([][]int, 0)

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		var fullLine []int
		line := scanner.Text()
		parts := strings.Split(line, ":")
		fullLine = append(fullLine, toInt(parts[0]))
		rightSide := strings.TrimSpace(parts[1])
		ints := strings.Split(rightSide, " ")
		for _, val := range ints {
			fullLine = append(fullLine, toInt(val))
		}
		res = append(res, fullLine)
	}

	return res
}

func toInt(val string) int {
	res, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return res
}
