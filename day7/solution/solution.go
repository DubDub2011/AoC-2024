package solution

func TotalCalibrationResult(input [][]int) int {
	total := 0
	for _, row := range input {
		targetValue := row[0]
		total += targetValue
	}
	return total
}
