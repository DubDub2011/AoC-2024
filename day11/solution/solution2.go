package solution

import (
	"strconv"
)

func StoneCountMapped(input []int, count int) int {
	res := make(map[int]int, 0)
	for _, val := range input {
		res[val]++
	}

	for idx := 0; idx < count; idx++ {
		res = applyRulesMap(res)
	}

	sum := 0
	for _, count := range res {
		sum += count
	}

	return sum
}

func applyRulesMap(input map[int]int) map[int]int {
	res := make(map[int]int, len(input))
	for val, count := range input {
		if val == 0 {
			res[1] += count
			continue
		}

		valStr := strconv.Itoa(val)
		if len(valStr)%2 == 0 {
			left, right := split(valStr)
			res[left] += count
			res[right] += count
			continue
		}

		val *= 2024
		res[val] += count
	}
	return res
}

func split(inp string) (int, int) {
	leftStr, rightStr := inp[:(len(inp)/2)], inp[(len(inp)/2):]
	left, err := strconv.Atoi(leftStr)
	if err != nil {
		panic(err)
	}
	right, err := strconv.Atoi(rightStr)
	if err != nil {
		panic(err)
	}

	return left, right
}

func runeToInt(val rune) int {
	if val < '0' || val > '9' {
		return 0
	}
	return int(val) - 48
}
