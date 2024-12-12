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
		{
			[][]rune{
				{'A', 'A'},
				{'A', 'A'},
			}, 32,
		},
		{
			[][]rune{
				{'A'},
				{'A', 'A', 'A'},
				{'A'},
			}, 60,
		},
		{
			[][]rune{
				{'A', 'B'},
				{'A', 'B'},
			}, 24,
		},
		{
			[][]rune{
				{'A', 'A', 'A', 'A'},
				{'B', 'B', 'C', 'D'},
				{'B', 'B', 'C', 'C'},
				{'E', 'E', 'E', 'C'},
			}, 140,
		},
		{
			[][]rune{
				{'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O'},
				{'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O'},
				{'O', 'O', 'O', 'O', 'O'},
			}, 772,
		},
		{
			[][]rune{
				{'R', 'R', 'R', 'R', 'I', 'I', 'C', 'C', 'F', 'F'},
				{'R', 'R', 'R', 'R', 'I', 'I', 'C', 'C', 'C', 'F'},
				{'V', 'V', 'R', 'R', 'R', 'C', 'C', 'F', 'F', 'F'},
				{'V', 'V', 'R', 'C', 'C', 'C', 'J', 'F', 'F', 'F'},
				{'V', 'V', 'V', 'V', 'C', 'J', 'J', 'C', 'F', 'E'},
				{'V', 'V', 'I', 'V', 'C', 'C', 'J', 'J', 'E', 'E'},
				{'V', 'V', 'I', 'I', 'I', 'C', 'J', 'J', 'E', 'E'},
				{'M', 'I', 'I', 'I', 'I', 'I', 'J', 'J', 'E', 'E'},
				{'M', 'I', 'I', 'I', 'S', 'I', 'J', 'E', 'E', 'E'},
				{'M', 'M', 'M', 'I', 'S', 'S', 'J', 'E', 'E', 'E'},
			}, 1930,
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

func TestFencePricerDiscount(t *testing.T) {
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
			}, 8,
		},
		{
			[][]rune{
				{'A', 'A'},
				{'A', 'A'},
			}, 16,
		},
		{
			[][]rune{
				{'A'},
				{'A', 'A', 'A'},
				{'A'},
			}, 40,
		},
		{
			[][]rune{
				{'A', 'B'},
				{'A', 'B'},
			}, 16,
		},
		{
			[][]rune{
				{'A', 'A', 'A', 'A'},
				{'B', 'B', 'C', 'D'},
				{'B', 'B', 'C', 'C'},
				{'E', 'E', 'E', 'C'},
			}, 80,
		},
		{
			[][]rune{
				{'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O'},
				{'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O'},
				{'O', 'O', 'O', 'O', 'O'},
			}, 436,
		},
		{
			[][]rune{
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'B', 'B', 'A'},
				{'A', 'A', 'A', 'B', 'B', 'A'},
				{'A', 'B', 'B', 'A', 'A', 'A'},
				{'A', 'B', 'B', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
			}, 368,
		},
		{
			[][]rune{
				{'R', 'R', 'R', 'R', 'I', 'I', 'C', 'C', 'F', 'F'},
				{'R', 'R', 'R', 'R', 'I', 'I', 'C', 'C', 'C', 'F'},
				{'V', 'V', 'R', 'R', 'R', 'C', 'C', 'F', 'F', 'F'},
				{'V', 'V', 'R', 'C', 'C', 'C', 'J', 'F', 'F', 'F'},
				{'V', 'V', 'V', 'V', 'C', 'J', 'J', 'C', 'F', 'E'},
				{'V', 'V', 'I', 'V', 'C', 'C', 'J', 'J', 'E', 'E'},
				{'V', 'V', 'I', 'I', 'I', 'C', 'J', 'J', 'E', 'E'},
				{'M', 'I', 'I', 'I', 'I', 'I', 'J', 'J', 'E', 'E'},
				{'M', 'I', 'I', 'I', 'S', 'I', 'J', 'E', 'E', 'E'},
				{'M', 'M', 'M', 'I', 'S', 'S', 'J', 'E', 'E', 'E'},
			}, 1206,
		},
		{
			[][]rune{
				{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'C', 'B', 'B', 'D', 'D', 'A'},
				{'A', 'A', 'C', 'B', 'B', 'A', 'A', 'A'},
				{'A', 'B', 'B', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'B', 'B', 'A', 'D', 'D', 'D', 'A'},
				{'A', 'A', 'A', 'A', 'D', 'A', 'D', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A', 'A', 'A'},
			}, 946,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			price := solution.FencePricerDiscount(tc.input)
			if price != tc.output {
				t.Errorf("expected %d, got %d", tc.output, price)
			}
		})
	}
}
