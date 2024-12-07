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
		{[][]int{{4, 1, 2}}, 0},
		{
			[][]int{
				{3, 1, 2},
				{2, 1, 2},
				{4, 1, 2},
			}, 5,
		},
		{
			[][]int{
				{190, 10, 19},
				{3267, 81, 40, 27},
				{83, 17, 5},
				{156, 15, 6},
				{7290, 6, 8, 6, 15},
				{161011, 16, 10, 13},
				{192, 17, 8, 14},
				{21037, 9, 7, 18, 13},
				{292, 11, 6, 16, 20},
			}, 3749,
		},
		{[][]int{{292, 11, 6, 16, 20}}, 292},
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

func TestTotalCalibrationResult2(t *testing.T) {
	testCases := []struct {
		input  [][]int
		output int
	}{
		{[][]int{{156, 15, 6}}, 156},
		{[][]int{{7290, 6, 8, 6, 15}}, 7290},
		{[][]int{{192, 17, 8, 14}}, 192},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.TotalCalibrationResult2(tc.input)

			if res != tc.output {
				t.Errorf("expected %d, got %d", tc.output, res)
			}
		})
	}
}
