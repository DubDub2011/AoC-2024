package main

import (
	"bufio"
	"day11/solution"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, count := readInput(), 25
	fmt.Printf("Part 1: %d\n", solution.StoneCount(data, count))
	count2 := 75
	fmt.Printf("Part 2: %d\n", solution.StoneCountSliced(data, count2))
}

func readInput() []int {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	res := make([]int, 0)
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		res = append(res, val)
	}

	return res
}
