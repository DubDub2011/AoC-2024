package solution_test

import (
	"day16/solution"
	"fmt"
	"testing"
)

func TestReindeerMaze(t *testing.T) {
	testCases := []struct {
		input  [][]rune
		output int
	}{
		{
			[][]rune{
				[]rune("#############"),
				[]rune("#S.........E#"),
				[]rune("#############"),
			}, 10,
		},
		{
			[][]rune{
				[]rune("###"),
				[]rune("#E#"),
				[]rune("#.#"),
				[]rune("#.#"),
				[]rune("#.#"),
				[]rune("#.#"),
				[]rune("#.#"),
				[]rune("#S#"),
				[]rune("###"),
			}, 1006,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.ReindeerMaze(tc.input)
			if res != tc.output {
				t.Errorf("Expected %d, got %d", tc.output, res)
			}
		})
	}

}
