package solution_test

import (
	"day8/solution"
	"fmt"
	"testing"
)

func TestResonance(t *testing.T) {
	testCases := []struct {
		input  [][]rune
		output int
	}{
		{
			[][]rune{
				{'.', '.', '.'},
				{'.', 'a', '.'},
				{'.', '.', '.'},
			}, 0,
		},
		{
			[][]rune{
				{'.', '.', '.'},
				{'.', 'a', '.'},
				{'.', '.', 'a'},
			}, 1,
		},
		{
			[][]rune{
				{'.', '.', '.', '.'},
				{'.', 'a', '.', '.'},
				{'.', '.', 'a', '.'},
				{'.', '.', '.', '.'},
			}, 2,
		},
		{
			[][]rune{
				{'.', '.', '.', '.'},
				{'.', 'a', '.', '.'},
				{'.', 'a', 'a', '.'},
				{'.', '.', '.', '.'},
			}, 6,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.Resonance(tc.input)

			if res != tc.output {
				t.Errorf("expected %d, got %d", tc.output, res)
			}
		})
	}
}
