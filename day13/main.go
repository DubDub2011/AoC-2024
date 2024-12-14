package main

import (
	"bufio"
	"day13/solution"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := readInput()
	fmt.Println("Part 1: %d", solution.ClawMachine(data))

}

func readInput() [][]int {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	res := make([][]int, 0)
	scanner := bufio.NewScanner(data)
	for {
		puzzle := []int{}
		scanner.Scan()
		a := scanner.Text()
		scanner.Scan()
		b := scanner.Text()
		scanner.Scan()
		total := scanner.Text()

		parts := strings.Split(a[len("Button A:"):], ",")
		for _, part := range parts {
			val, err := strconv.Atoi(part[len(" X+"):])
			if err != nil {
				panic(err)
			}
			puzzle = append(puzzle, val)
		}

		parts = strings.Split(b[len("Button B:"):], ",")
		for _, part := range parts {
			val, err := strconv.Atoi(part[len(" X+"):])
			if err != nil {
				panic(err)
			}
			puzzle = append(puzzle, val)
		}

		parts = strings.Split(total[len("Prize:"):], ",")
		for _, part := range parts {
			val, err := strconv.Atoi(part[len(" X="):])
			if err != nil {
				panic(err)
			}
			puzzle = append(puzzle, val)
		}

		if len(puzzle) != 6 {
			panic(fmt.Errorf("bad input %v", puzzle))
		}

		res = append(res, puzzle)
		if !scanner.Scan() {
			break
		}
	}

	return res
}
