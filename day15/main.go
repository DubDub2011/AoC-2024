package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mapData, commands := readInput()
	fmt.Printf("Map %v\n", mapData)
	fmt.Printf("Commands %s\n", string(commands))
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
