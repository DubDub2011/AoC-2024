package solution

import (
	"fmt"
	"math"
	"strconv"
)

func StoneCountSliced(input []int, count int) int {
	res := make([][]int, 0, len(input))
	for _, val := range input {
		res = append(res, buildIntSlice(val))
	}

	for idx := 0; idx < count; idx++ {
		res = applyRulesSliced(res)
		fmt.Printf("Length at %d is: %d\n", idx, len(res))
	}

	return len(res)
}

func applyRulesSliced(input [][]int) [][]int {
	res := make([][]int, 0, len(input))
	for _, intSlice := range input {
		if len(intSlice) == 1 && intSlice[0] == 0 {
			res = append(res, []int{1})
			continue
		}

		if len(intSlice)%2 == 0 {
			left, right := splitSlice(intSlice)
			res = append(res, left)
			res = append(res, right)
			continue
		}

		val := buildInt(intSlice)
		val *= 2024
		res = append(res, buildIntSlice(val))
	}
	return res
}

func splitSlice(inp []int) ([]int, []int) {
	left, right := inp[:(len(inp)/2)], inp[(len(inp)/2):]
	left = buildIntSlice(buildInt((left)))   // needed to handle 0s
	right = buildIntSlice(buildInt((right))) // needed to handle 0s

	return left, right
}

func runeToInt(val rune) int {
	if val < '0' || val > '9' {
		return 0
	}
	return int(val) - 48
}

func buildInt(digits []int) int {
	result := 0
	length := len(digits)
	for i, digit := range digits {
		if digit < 0 || digit > 9 {
			panic("invalid digit")
		}

		mag := int(math.Pow10(length - i - 1))
		result += digit * mag
	}
	return result
}

func buildIntSlice(val int) []int {
	strVal := strconv.Itoa(val)
	res := make([]int, len(strVal))
	for idx, digit := range strVal {
		res[idx] = runeToInt(digit)
	}
	return res
}
