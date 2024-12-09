package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Input: %s", readInput())
}

func readInput() []byte {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	res := make([]byte, 0)
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		res = append(res, []byte(scanner.Text())...)
	}

	return res
}
