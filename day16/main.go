package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// looks like it's time to learn A*
	mapData := readInput()
	fmt.Printf("Map data: %v\n", mapData)
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
