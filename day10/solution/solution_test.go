package solution_test

import (
	"day10/solution"
	"fmt"
	"testing"
)

func TestHeightmapRouter(t *testing.T) {
	testCases := []struct {
		input  [][]rune
		output int
	}{
		{
			[][]rune{
				{'0', '.', '.', '.', '.', '.', '.'},
				{'1', '.', '.', '.', '.', '.', '.'},
				{'2', '.', '.', '.', '.', '.', '.'},
				{'3', '.', '.', '.', '.', '.', '.'},
				{'4', '.', '.', '.', '.', '.', '.'},
				{'5', '.', '.', '.', '.', '.', '.'},
				{'6', '7', '8', '9', '.', '.', '.'},
			}, 1,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.HeightmapRouter(tc.input)
			if res != tc.output {
				t.Errorf("expected %d, got %d", tc.output, res)
			}
		})

	}
}
