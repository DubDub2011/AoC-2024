package main

import (
	"bufio"
	"day14/solution"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readInput()
	fmt.Printf("Part 1: %d\n", solution.RobotPredictor(data, 100))
	fmt.Printf("Part 2: %d\n", solution.RobotChristmasTree(data))
}

func readInput() [][]int {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	res := make([][]int, 0)
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		positionLine := parts[0]
		coords := strings.Split(positionLine[len("p="):], ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			panic(err)
		}

		velocityLine := parts[1]
		velocity := strings.Split(velocityLine[len("v="):], ",")
		vx, err := strconv.Atoi(velocity[0])
		if err != nil {
			panic(err)
		}
		vy, err := strconv.Atoi(velocity[1])
		if err != nil {
			panic(err)
		}

		res = append(res, []int{x, y, vx, vy})
	}

	return res
}
