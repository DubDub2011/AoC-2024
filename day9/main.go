package main

import (
	"bufio"
	"day9/solution"
	"fmt"
	"os"
)

func main() {
	data := readInput()
	fmt.Printf("Part 1: %d\n", solution.DiskFragmenter(data))
	data = readInput()
	fmt.Printf("Part 2: %d\n", solution.DiskFragmenter2(data)) // still wrong but no clue why...
}

func readInput() []byte {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	res := make([]byte, 0)
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		res = append(res, []byte(scanner.Text())...)
	}

	return res
}
