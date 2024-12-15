package solution

import (
	gr "day15/grid"
	"fmt"
	"testing"
)

func TestValidMovePositions(t *testing.T) {
	testCases := []struct {
		input   [][]rune
		moveDir func(gr.Position) gr.Position
		output  []gr.Position
	}{
		{
			[][]rune{
				[]rune("#####"),
				[]rune("#...#"),
				[]rune("#.@.#"),
				[]rune("#...#"),
				[]rune("#####"),
			},
			gr.Position.Right,
			[]gr.Position{gr.Position{2, 2}},
		},
		{
			[][]rune{
				[]rune("#####"),
				[]rune("#...#"),
				[]rune("#.@O#"),
				[]rune("#...#"),
				[]rune("#####"),
			},
			gr.Position.Right,
			[]gr.Position{},
		},
		{
			[][]rune{
				[]rune("######"),
				[]rune("#....#"),
				[]rune("#.@O.#"),
				[]rune("#....#"),
				[]rune("#....#"),
				[]rune("######"),
			},
			gr.Position.Right,
			[]gr.Position{gr.Position{2, 2}, gr.Position{3, 2}},
		},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			grid := gr.New(testCase.input)

			actual := getValidMovePositions(gr.Position{2, 2}, grid, testCase.moveDir)
			if len(actual) != len(testCase.output) {
				t.Errorf("Expected %v, got %v", testCase.output, actual)
			}

			for idx, pos := range testCase.output {
				if pos != actual[idx] {
					t.Errorf("Expected %v, got %v", testCase.output, actual)
				}
			}
		})
	}
}
