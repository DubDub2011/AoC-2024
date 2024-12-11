package solution_test

import (
	"day11/solution"
	"fmt"
	"testing"
)

func TestStoneCount(t *testing.T) {
	testCases := []struct {
		input  []int
		count  int
		output int
	}{
		{[]int{0}, 1, 1}, // 1
		{[]int{0}, 2, 1}, // 2024
		{[]int{0}, 3, 2}, // 20 24
		{[]int{0}, 4, 4}, // 2 0 2 4
		{[]int{0}, 5, 4}, // 4048 0 4048 8096
		{[]int{0}, 6, 7}, // 40 48 0 40 48 80 96
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.StoneCount(tc.input, tc.count)
			if res != tc.output {
				t.Errorf("Expected %d, got %d", tc.output, res)
			}
		})
	}
}
