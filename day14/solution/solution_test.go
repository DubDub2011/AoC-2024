package solution_test

import (
	"day14/solution"
	"fmt"
	"testing"
)

func TestRobotPredictor(t *testing.T) {
	testCases := []struct {
		robots  [][]int
		seconds int
		output  int
	}{
		{
			[][]int{
				{0, 0, 1, 1},
			}, 1, 0,
		},
		{
			[][]int{
				{0, 0, 1, 1},
				{100, 0, -1, 1},
				{0, 100, 1, -1},
				{100, 100, -1, 1},
			}, 1, 1,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.RobotPredictor(tc.robots, tc.seconds)
			if res != tc.output {
				t.Errorf("Expected %d, got %d", tc.output, res)
			}
		})
	}
}
