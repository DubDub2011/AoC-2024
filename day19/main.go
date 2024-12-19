package main

import (
	"bufio"
	"day19/solution"
	"fmt"
	"os"
	"strings"
)

func main() {
	patterns, target := readInput()
	fmt.Printf("Part 1: %d\n", solution.PossibleDesigns(patterns, target))
	fmt.Printf("Part 2: %d\n", solution.AllDesigns(patterns, target))
}

func readInput() ([]string, []string) {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)
	scanner.Scan()

	availablePatterns := strings.Split(scanner.Text(), ", ")

	scanner.Scan() // empty line

	targetDesigns := []string{}
	for scanner.Scan() {
		targetDesigns = append(targetDesigns, scanner.Text())
	}

	return availablePatterns, targetDesigns
}
