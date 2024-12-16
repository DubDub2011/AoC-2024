package main

import (
	"bufio"
	"day16/solution"
	"fmt"
	"os"
)

func main() {
	mapData := readInput()
	fmt.Printf("Part 1: %d\n", solution.ReindeerMaze(mapData))
}

func readInput() [][]rune {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	mapData := make([][]rune, 0)
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		mapData = append(mapData, []rune(scanner.Text()))
	}

	return mapData
}
