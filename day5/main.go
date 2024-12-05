package main

import (
	"bufio"
	"day5/solution"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pairs, order := readInput()
	fmt.Printf("Part 1: %d\n", solution.PagePrint(pairs, order))
	fmt.Printf("Part 2: %d\n", solution.PagePrint2(pairs, order))
}

func readInput() ([][]int, [][]int) {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)
	pairs := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break // first part done
		}

		parts := strings.Split(line, "|")
		left, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		pairs = append(pairs, []int{left, right})
	}

	order := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break // first part done
		}

		row := make([]int, 0)
		parts := strings.Split(line, ",")
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			row = append(row, num)
		}
		order = append(order, row)
	}

	return pairs, order
}
