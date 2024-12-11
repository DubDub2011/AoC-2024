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
		{[]int{0}, 1, 1},
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
