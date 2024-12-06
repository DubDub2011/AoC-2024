package solution_test

import (
	"fmt"
	"testing"

	"day6/solution"
)

func TestRouteFinder(t *testing.T) {
	testCases := []struct {
		input  [][]rune
		output int
	}{
		{
			[][]rune{
				{'.', '#', '.'},
				{'.', '.', '.'},
				{'.', '^', '.'},
			}, 3,
		},
		{
			[][]rune{
				{'.', '#', '.', '.'},
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
				{'.', '^', '.', '.'},
			}, 5,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			result, _ := solution.RouteFinder(tc.input)
			if result != tc.output {
				t.Errorf("Case failed, expected %d, got %d", tc.output, result)
			}
		})
	}
}

func TestRouteBreaker(t *testing.T) {
	testCases := []struct {
		input  [][]rune
		output int
	}{
		{
			[][]rune{
				{'.', '#', '.', '.'},
				{'.', '.', '.', '#'},
				{'.', '.', '.', '.'},
				{'.', '^', '#', '.'},
			}, 1,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			result := solution.RouteBreaker(tc.input)
			if result != tc.output {
				t.Errorf("Case failed, expected %d, got %d", tc.output, result)
			}
		})
	}
}
