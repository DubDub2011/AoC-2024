package solution_test

import (
	"day12/solution"
	"fmt"
	"testing"
)

func TestFencePricer(t *testing.T) {
	testCases := []struct {
		input  [][]rune
		output int
	}{
		{
			[][]rune{
				{'A'},
			}, 4,
		},
		{
			[][]rune{
				{'A', 'A'},
			}, 12,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			price := solution.FencePricer(tc.input)
			if price != tc.output {
				t.Errorf("expected %d, got %d", tc.output, price)
			}
		})
	}
}
