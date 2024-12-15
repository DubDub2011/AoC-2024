package main

import (
	"bufio"
	"day15/grid"
	"day15/solution"
	"fmt"
	"os"
)

func main() {
	mapData, commands := readInput()
	fmt.Printf("Part 1 %d\n", solution.BoxChecker(mapData, commands))

	mapData, commands = readInputBigBox()
	grid := grid.New(mapData)
	fmt.Printf("Starting grid: %s\n", grid.String())
	fmt.Printf("Part 2 %d\n", solution.BigBoxChecker(mapData, commands)) // was going to test more but this just worked

}

func readInput() ([][]rune, []rune) {
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

	commands := make([]rune, 0)
	for scanner.Scan() {
		commands = append(commands, []rune(scanner.Text())...)
	}

	return mapData, commands

}

func readInputBigBox() ([][]rune, []rune) {
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
		line := scanner.Text()
		res := make([]rune, 0)
		for _, c := range line {
			if c == '#' {
				res = append(res, '#')
				res = append(res, '#')
			} else if c == '.' {
				res = append(res, '.')
				res = append(res, '.')
			} else if c == 'O' {
				res = append(res, '[')
				res = append(res, ']')
			} else if c == '@' {
				res = append(res, '@')
				res = append(res, '.')
			}
		}

		mapData = append(mapData, res)
	}

	commands := make([]rune, 0)
	for scanner.Scan() {
		commands = append(commands, []rune(scanner.Text())...)
	}

	return mapData, commands

}
