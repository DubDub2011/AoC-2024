package main

import (
	"bufio"
	"fmt"
	"os"

	"day4/solution"
)

func main() {
	data := readInput()
	fmt.Printf("Part 1: %d\n", solution.XMASOccurances(data))
	fmt.Printf("Part 2: %d\n", solution.XMASOccurances2(data))
}

func readInput() [][]rune {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	res := make([][]rune, 0)
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		res = append(res, []rune(scanner.Text()))
	}

	return res
}
