package main

import (
	"day3/solution"
	"fmt"
	"os"
)

func main() {
	data := readInput()
	fmt.Printf("Part 1: %d\n", solution.Multiplier(string(data)))
	fmt.Printf("Part 2: %d\n", solution.Multiplier2(string(data)))
}

func readInput() []byte {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return data
}
