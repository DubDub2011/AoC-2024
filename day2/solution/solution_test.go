package solution_test

import (
	"day2/solution"
	"fmt"
	"testing"
)

func TestReportSafeCount(t *testing.T) {
	testCases := []struct {
		input    [][]int
		expected int
	}{
		{[][]int{{1, 2, 3, 4}}, 1},
		{[][]int{{7, 6, 4, 2, 1}}, 1},
		{[][]int{{1, 2, 7, 8, 9}}, 0},
		{[][]int{{9, 7, 6, 2, 1}}, 0},
		{[][]int{{1, 3, 2, 4, 5}}, 0},
		{[][]int{{8, 6, 4, 4, 1}}, 0},
		{[][]int{{1, 3, 6, 7, 9}}, 1},
		{
			[][]int{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			}, 2,
		},
	}

	for idx, test := range testCases {
		t.Run(fmt.Sprintf("Case %d", idx), func(t *testing.T) {
			res := solution.ReportSafeCount(test.input)

			if res != test.expected {
				t.Errorf("expected %d, got %d", test.expected, res)
			}
		})
	}

}

func TestReportSafeCount2(t *testing.T) {
	testCases := []struct {
		input    [][]int
		expected int
	}{
		{[][]int{{1, 2, 3, 4}}, 1},
		{[][]int{{7, 6, 4, 2, 1}}, 1},
		{[][]int{{1, 2, 7, 8, 9}}, 0},
		{[][]int{{9, 7, 6, 2, 1}}, 0},
		{[][]int{{1, 3, 2, 4, 5}}, 1},
		{[][]int{{8, 6, 4, 4, 1}}, 1},
		{[][]int{{1, 3, 6, 7, 9}}, 1},
		{
			[][]int{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			}, 4,
		},
	}

	for idx, test := range testCases {
		t.Run(fmt.Sprintf("Case %d", idx), func(t *testing.T) {
			res := solution.ReportSafeCount2(test.input)

			if res != test.expected {
				t.Errorf("expected %d, got %d", test.expected, res)
			}
		})
	}

}
