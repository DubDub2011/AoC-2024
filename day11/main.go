package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Input: ", readInput())
}

func readInput() []int {
	data, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	res := make([]int, 0)
	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		res = append(res, val)
	}

	return res
}
