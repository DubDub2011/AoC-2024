package solution

import "fmt"

func ClawMachine(input [][]int) int {
	// wonder if this wont work because computers don't math
	tokens := 0
	for _, puzzle := range input {
		A := [][]int{
			{puzzle[0], puzzle[1]},
			{puzzle[2], puzzle[3]},
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
			(inverseA[0][0]*B[0] + inverseA[0][1]*B[1]) / determinant,
			(inverseA[1][0]*B[0] + inverseA[1][1]*B[1]) / determinant,
		}

		fmt.Printf("Result %v\n", result)

		tokens += result[0]*3 + result[1]
	}
	return tokens
}
