package main

import (
	"bufio"
	"day17/solution"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	registers, instructions := readInput()
	fmt.Printf("Registers %v\n", registers)
	fmt.Printf("Instructions %v\n", instructions)
	fmt.Printf("Part 1: %s\n", solution.ThreeBitComputer(registers, instructions))
	solution.Repeat(registers, instructions)
}

func readInput() ([]int, []int) {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	registers := make([]int, 0, 3)
	scanner := bufio.NewScanner(data)
	scanner.Scan()
	line := scanner.Text()[len("Register A: "):]
	toInt, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	registers = append(registers, toInt)

	scanner.Scan()
	line = scanner.Text()[len("Register B: "):]
	toInt, err = strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	registers = append(registers, toInt)

	scanner.Scan()
	line = scanner.Text()[len("Register C: "):]
	toInt, err = strconv.Atoi(line)
	if err != nil {
		panic(err)
	}
	registers = append(registers, toInt)

	scanner.Scan() // empty line

	scanner.Scan()
	line = scanner.Text()[len("Program: "):]
	instructions := make([]int, 0)
	for _, c := range strings.Split(line, ",") {
		toInt, err = strconv.Atoi(c)
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, toInt)
	}
	return registers, instructions
}
