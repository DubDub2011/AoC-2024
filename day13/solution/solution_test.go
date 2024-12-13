package solution_test

import (
	"day13/solution"
	"fmt"
	"testing"
)

func TestClawMachine(t *testing.T) {
	testCases := []struct {
		input  [][]float64
		output int
	}{
		{[][]float64{
			{3, 2, 5, 7, 23, 31}, // this doesn't work, infinite solutions
		}, 7},
		{[][]float64{
			{94, 34, 22, 67, 8400, 5400},
		}, 280},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.ClawMachine(tc.input)
			if res != tc.output {
				t.Errorf("Expected %d, got %d", tc.output, res)
			}
		})
	}
}
