package main

import (
	"bufio"
	"day8/solution"
	"fmt"
	"os"
)

func main() {
	data := readInput()
	fmt.Printf("Part 1: %d", solution.Resonance(data))
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
