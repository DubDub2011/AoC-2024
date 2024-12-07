package solution

import (
	"strconv"
)

func TotalCalibrationResult2(input [][]int) int {
	total := 0
	for _, row := range input {
		targetValue := row[0]
		row = row[1:]
		if !possible2(targetValue, row) {
			continue
		}
		total += targetValue
	}
	return total
}

func possible2(target int, input []int) bool {
	limit := pow(3, (len(input) - 1))
	for operatorVal := 0; operatorVal < limit; operatorVal++ {
		// bit tricky but we can use bits to determine the operators
		res := input[0]
		for position := range input {
			if position == len(input)-1 {
				break
			}

			posVal := (operatorVal / pow(3, position)) % 3
			if posVal == 0 {
				res += input[position+1]
			} else if posVal == 1 {
				res *= input[position+1]
			} else {
				temp := strconv.Itoa(res) + strconv.Itoa(input[position+1])
				var err error
				res, err = strconv.Atoi(temp)
				if err != nil {
					panic(err)
				}
			}
		}

		if res == target {
			return true
		}
	}
	return false
}
