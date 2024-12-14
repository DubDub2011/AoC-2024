package solution_test

import (
	"day13/solution"
	"fmt"
	"testing"
)

func TestClawMachine(t *testing.T) {
	testCases := []struct {
		input  [][]int
		output int
	}{
		{[][]int{
			{1, 1, 3, 2, 6, 12},
		}, 6},
		{[][]int{
			{2, 1, -1, 1, 5, 2},
		}, 6},
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
