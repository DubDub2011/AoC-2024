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
			{94, 34, 22, 67, 8400, 5400},
		}, 280},
		{[][]int{
			{26, 66, 67, 21, 12748, 12176},
		}, 0},
		{[][]int{
			{17, 86, 84, 37, 7870, 6450},
		}, 200},
		{[][]int{
			{69, 23, 27, 71, 18641, 10279},
		}, 0},
		{[][]int{
			{94, 34, 22, 67, 8400, 5400},
			{26, 66, 67, 21, 12748, 12176},
			{17, 86, 84, 37, 7870, 6450},
			{69, 23, 27, 71, 18641, 10279},
		}, 480},
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
