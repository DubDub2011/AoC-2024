package solution_test

import (
	"day15/solution"
	"fmt"
	"testing"
)

func TestRobotMover(t *testing.T) {

	testCases := []struct {
		mapData  [][]rune
		commands []rune
		output   int
	}{
		{
			[][]rune{
				[]rune("#####"),
				[]rune("#...#"),
				[]rune("#.@.#"),
				[]rune("#...#"),
				[]rune("#####"),
			}, []rune("<v>>^^<<v>"), 0,
		},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.RobotMover(testCase.mapData, testCase.commands)
			if res != testCase.output {
				t.Errorf("Expected %d, got %d", testCase.output, res)
			}

		})
	}
}
