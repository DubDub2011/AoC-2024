package solution_test

import (
	"day7/solution"
	"fmt"
	"testing"
)

func TestTotalCalibrationResult(t *testing.T) {
	testCases := []struct {
		input  [][]int
		output int
	}{
		{[][]int{{3, 1, 2}}, 3},
		{[][]int{{2, 1, 2}}, 2},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.TotalCalibrationResult(tc.input)

			if res != tc.output {
				t.Errorf("expected %d, got %d", tc.output, res)
			}
		})
	}
}
