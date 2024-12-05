package solution

import (
	"fmt"
)

func ReportSafeCount(input [][]int) int {
	total := 0

	for _, row := range input {
		safe := true
		if len(row) < 2 {
			panic(fmt.Errorf("row of length less than 2 found, %v", row))
		}

		prevValue := row[0]
		ascending := row[0] < row[1]
		for _, val := range row[1:] {
			diff := prevValue - val
			if diff == 0 {
				safe = false
				break
			}

			if diff > 3 || diff < -3 {
				safe = false
				break
			}

			if ascending && diff > 0 {
				safe = false
				break
			} else if !ascending && diff < 0 {
				safe = false
				break
			}
			prevValue = val
		}

		if safe {
			total++
		}
	}
	return total
}

func ReportSafeCount2(input [][]int) int {
	total := 0

	for _, parentRow := range input {
		for idx := 0; idx < len(parentRow); idx++ {
			row := make([]int, len(parentRow))
			copy(row, parentRow)
			row = append(row[:idx], row[idx+1:]...)
			if checkReport(row) {
				total++
				break
			}
		}
	}
	return total
}

func checkReport(input []int) bool {
	prevValue := input[0]
	ascending := input[0] < input[1]
	for _, val := range input[1:] {
		diff := prevValue - val
		if diff == 0 {
			return false
		}

		if diff > 3 || diff < -3 {
			return false
		}

		if ascending && diff > 0 {
			return false
		} else if !ascending && diff < 0 {
			return false
		}
		prevValue = val
	}
	return true
}
