package solution

import (
	gr "day14/grid"
	"fmt"
	"testing"
)

func TestRobot(t *testing.T) {
	testCases := []struct {
		robo    Robot
		seconds int
		output  gr.Position
	}{
		{Robot{gr.Position{0, 0}, 0, 0}, 1, gr.Position{0, 0}},
		{Robot{gr.Position{0, 0}, 1, 1}, 1, gr.Position{1, 1}},
		{Robot{gr.Position{0, 0}, 1, 1}, 50, gr.Position{50, 50}},
		{Robot{gr.Position{0, 0}, -1, -1}, 1, gr.Position{100, 102}},
		{Robot{gr.Position{0, 0}, -1, -1}, 50, gr.Position{51, 53}},
		{Robot{gr.Position{0, 0}, 1, 1}, 10, gr.Position{10, 10}},
		{Robot{gr.Position{0, 0}, 1, 1}, 150, gr.Position{49, 47}},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			testCase.robo.Move(testCase.seconds)
			if testCase.robo.Position() != testCase.output {
				t.Errorf("Expected %v, got %v", testCase.output, testCase.robo.Position())
			}
		})
	}
}
