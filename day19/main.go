package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	patterns, target := readInput()
	fmt.Printf("Patterns: %v\n", patterns)
	fmt.Printf("Target: %v\n", target)
}

func readInput() ([]string, []string) {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(data)
	scanner.Scan()

	availablePatterns := strings.Split(scanner.Text(), ",")

	scanner.Scan() // empty line

	targetDesigns := []string{}
	for scanner.Scan() {
		targetDesigns = append(targetDesigns, scanner.Text())
	}

	return availablePatterns, targetDesigns
}
