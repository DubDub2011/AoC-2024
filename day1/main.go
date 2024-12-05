package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	left, right := readInput()

	part1(left, right)
	part2(left, right)
}

func part1(left, right []int) {
	sort.Ints(left)
	sort.Ints(right)

	idx := 0
	total := 0
	for {
		if idx == len(left) {
			break
		}

		leftSide := left[idx]
		rightSide := right[idx]
		if leftSide > rightSide {
			total += leftSide - rightSide
		} else {
			total += rightSide - leftSide
		}
		idx++
	}

	fmt.Printf("Part 1: %d\n", total)
}

func part2(left, right []int) {
	mappingTable := make(map[int]int)
	for _, occurence := range right {
		mappingTable[occurence]++
	}

	total := 0
	for _, val := range left {
		total += val * mappingTable[val]
	}

	fmt.Printf("Part 2: %d\n", total)
}

func readInput() ([]int, []int) {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	var leftSide, rightSide []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		left, err := strconv.Atoi(line[:5])
		if err != nil {
			panic(err)
		}

		right, err := strconv.Atoi(line[8:])
		if err != nil {
			panic(err)
		}

		leftSide = append(leftSide, left)
		rightSide = append(rightSide, right)
	}

	return leftSide, rightSide
}
