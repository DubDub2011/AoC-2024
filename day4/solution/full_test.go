package solution_test

import (
	"fmt"
	"testing"

	"day4/solution"
)

func TestXMASOccurances(t *testing.T) {
	testCases := []struct {
		input  [][]rune
		output int
	}{
		{
			[][]rune{
				{'X', 'M', 'A', 'S'},
			}, 1,
		},
		{
			[][]rune{
				{'X', 'M', 'A', 'S'},
				{'X', 'M', 'A', 'S'},
				{'X', 'M', 'A', 'S'},
			}, 3,
		},
		{
			[][]rune{
				{'X', 'X', 'A', 'S'},
				{'X', 'M', 'A', 'S'},
				{'X', 'M', 'A', 'S'},
			}, 2,
		},
		{
			[][]rune{
				{'X', 'M', 'A', 'S'},
				{'S', 'A', 'M', 'X'},
				{'X', 'M', 'A', 'S'},
			}, 3,
		},
		{
			[][]rune{
				{'X', 'S', 'X', 'S'},
				{'M', 'A', 'M', 'A'},
				{'A', 'M', 'A', 'M'},
				{'S', 'X', 'S', 'X'},
			}, 4,
		},
		{
			[][]rune{
				{'X', 'S', 'X', 'X'},
				{'M', 'M', 'M', 'M'},
				{'A', 'A', 'A', 'A'},
				{'S', 'X', 'S', 'S'},
			}, 5,
		},
	}

	for idx, testCase := range testCases {
		t.Log(fmt.Sprintf("Test case %d", idx))
		actual := solution.XMASOccurances(testCase.input)
		if actual != testCase.output {
			t.Errorf("XMASOccurances(%v) = %v, expected %v", testCase.input, actual, testCase.output)
		}
	}

}

func TestXMASOccurances2(t *testing.T) {
	testCases := []struct {
		input  [][]rune
		output int
	}{
		{
			[][]rune{
				{'X', 'M', 'A', 'S'},
			}, 0,
		},
		{
			[][]rune{
				{'X', 'M', 'A', 'S'},
				{'X', 'M', 'A', 'S'},
				{'X', 'M', 'A', 'S'},
			}, 1,
		},
		{
			[][]rune{
				{'X', 'X', 'A', 'S'},
				{'X', 'M', 'A', 'S'},
				{'X', 'M', 'A', 'S'},
			}, 0,
		},
		{
			[][]rune{
				{'X', 'M', 'A', 'S'},
				{'S', 'A', 'M', 'X'},
				{'X', 'M', 'A', 'S'},
			}, 0,
		},
		{
			[][]rune{
				{'X', 'S', 'X', 'S'},
				{'M', 'A', 'M', 'A'},
				{'A', 'M', 'A', 'M'},
				{'S', 'X', 'S', 'X'},
			}, 0,
		},
		{
			[][]rune{
				{'X', 'S', 'X', 'X'},
				{'M', 'M', 'M', 'M'},
				{'A', 'A', 'A', 'A'},
				{'S', 'X', 'S', 'S'},
			}, 1,
		},
	}

	for idx, testCase := range testCases {
		t.Log(fmt.Sprintf("Test case %d", idx))
		actual := solution.XMASOccurances2(testCase.input)
		if actual != testCase.output {
			t.Errorf("XMASOccurances(%v) = %v, expected %v", testCase.input, actual, testCase.output)
		}
	}

}
