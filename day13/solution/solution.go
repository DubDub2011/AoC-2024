package solution

import "fmt"

func ClawMachine(input [][]int) int {
	tokens := 0
	for _, puzzle := range input {
		A := [][]int{
			{puzzle[0], puzzle[2]},
			{puzzle[1], puzzle[3]},
		}

		B := []int{
			puzzle[4],
			puzzle[5],
		}

		// Use matrix method
		determinant := A[0][0]*A[1][1] - A[0][1]*A[1][0]
		if determinant == 0 {
			continue
		}

		inverseA := [][]int{
			{A[1][1], -A[0][1]},
			{-A[1][0], A[0][0]},
		}

		result := []int{
			(inverseA[0][0]*B[0] + inverseA[0][1]*B[1]),
			(inverseA[1][0]*B[0] + inverseA[1][1]*B[1]),
		}

		if result[0]%determinant != 0 || result[1]%determinant != 0 {
			continue
		}

		result = []int{
			result[0] / determinant,
			result[1] / determinant,
		}

		if result[0] < 0 || result[1] < 0 {
			continue
		}

		fmt.Printf("Result %v\n", result)

		tokens += result[0]*3 + result[1]
	}
	return tokens
}
