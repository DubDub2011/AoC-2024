package main

import (
	"bufio"
	"day6/solution"
	"fmt"
	"os"
)

func main() {
	points := readInput()
	routeLength, _ := solution.RouteFinder(points)
	fmt.Printf("Part 1: %d", routeLength)
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
