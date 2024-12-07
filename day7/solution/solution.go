package solution

import (
	"math"
)

func TotalCalibrationResult(input [][]int) int {
	total := 0
	for _, row := range input {
		targetValue := row[0]
		row = row[1:]
		if !possible(targetValue, row) {
			continue
		}
		total += targetValue
	}
	return total
}

func possible(target int, input []int) bool {
	limit := pow(2, (len(input) - 1))
	for operatorVal := 0; operatorVal < limit; operatorVal++ {
		// bit tricky but we can use bits to determine the operators
		operators := operatorVal

		res := input[0]
		for position := range input {
			if position == len(input)-1 {
				break
			}

			if (operators>>position)&1 == 0 {
				res += input[position+1]
			} else {
				res *= input[position+1]
			}
		}

		if res == target {
			return true
		}
	}
	return false
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
