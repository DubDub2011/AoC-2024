package solution_test

import (
	"day5/solution"
	"fmt"
	"testing"
)

func TestPagePrint(t *testing.T) {
	testCases := []struct {
		order  [][]int
		pages  [][]int
		output int
	}{
		{
			[][]int{{47, 53}, {50, 53}},
			[][]int{{47, 50, 53}},
			50,
		},
		{
			[][]int{{47, 53}, {50, 53}},
			[][]int{{47, 53, 50}},
			0,
		},
		{
			[][]int{
				{47, 53},
				{97, 13},
				{97, 61},
				{97, 47},
				{75, 29},
				{61, 13},
				{75, 53},
				{29, 13},
				{97, 29},
				{53, 29},
				{61, 53},
				{97, 53},
				{61, 29},
				{47, 13},
				{75, 47},
				{97, 75},
				{47, 61},
				{75, 61},
				{47, 29},
				{75, 13},
				{53, 13},
			},
			[][]int{
				{75, 47, 61, 53, 29},
				{97, 61, 53, 29, 13},
				{75, 29, 13},
				{75, 97, 47, 61, 53},
				{61, 13, 29},
				{97, 13, 75, 29, 47},
			},
			143,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", idx), func(t *testing.T) {
			output := solution.PagePrint(tc.order, tc.pages)

			if output != tc.output {
				t.Errorf("expected %d, got %d", tc.output, output)
			}
		})
	}
}

func TestPagePrint2(t *testing.T) {
	testCases := []struct {
		order  [][]int
		pages  [][]int
		output int
	}{
		{
			[][]int{{47, 53}, {50, 53}},
			[][]int{{47, 50, 53}},
			0,
		},
		{
			[][]int{{47, 50}, {50, 53}},
			[][]int{{47, 53, 50}},
			50,
		},
		{
			[][]int{
				{47, 53},
				{97, 13},
				{97, 61},
				{97, 47},
				{75, 29},
				{61, 13},
				{75, 53},
				{29, 13},
				{97, 29},
				{53, 29},
				{61, 53},
				{97, 53},
				{61, 29},
				{47, 13},
				{75, 47},
				{97, 75},
				{47, 61},
				{75, 61},
				{47, 29},
				{75, 13},
				{53, 13},
			},
			[][]int{
				{75, 47, 61, 53, 29},
				{97, 61, 53, 29, 13},
				{75, 29, 13},
				{75, 97, 47, 61, 53},
				{61, 13, 29},
				{97, 13, 75, 29, 47},
			},
			123,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", idx), func(t *testing.T) {
			output := solution.PagePrint2(tc.order, tc.pages)

			if output != tc.output {
				t.Errorf("expected %d, got %d", tc.output, output)
			}
		})
	}
}
