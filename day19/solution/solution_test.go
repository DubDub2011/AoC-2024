package solution_test

import (
	"day19/solution"
	"fmt"
	"testing"
)

func TestPossibleDesigns(t *testing.T) {
	testCases := []struct {
		availablePatterns []string
		targetPatterns    []string
		possibleDesigns   int
	}{
		{
			[]string{"a"},
			[]string{"aaa"},
			1,
		},
		{
			[]string{"a", "b"},
			[]string{"aba"},
			1,
		},
		{
			[]string{"a", "b"},
			[]string{"abc"},
			0,
		},
		{
			[]string{"a", "aa", "ab", "b"},
			[]string{"aaabbaba"},
			1,
		},
		{
			[]string{"a", "aa", "ab", "b"},
			[]string{"aaabbaba"},
			1,
		},
		{
			[]string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"},
			[]string{
				"brwrr",
				"bggr",
				"gbbr",
				"rrbgbr",
				"ubwu",
				"bwurrg",
				"brgr",
				"bbrgwb",
			},
			6,
		},
		{
			[]string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"},
			[]string{
				"brwrrbrwrrbrwrrbrwrrbrwrrbrwrrbrwrrbrwrrbrwrrbrwrrbrwrrbrwrrbrwrrbrwrrbrwrrbrwrr",
				"bggrbggrbggrbggrbggrbggrbggrbggrbggrbggrbggrbggrbggrbggrbggrbggrbggrbggrbggrbggr",
				"gbbrgbbrgbbrgbbrgbbrgbbrgbbrgbbrgbbrgbbrgbbrgbbrgbbrgbbrgbbrgbbrgbbrgbbrgbbrgbbr",
				"rrbgbrrrbgbrrrbgbrrrbgbrrrbgbrrrbgbrrrbgbrrrbgbrrrbgbrrrbgbrrrbgbrrrbgbrrrbgbrrrbgbr",
				"ubwuubwuubwuubwuubwuubwuubwuubwuubwuubwuubwuubwuubwuubwuubwuubwuubwuubwuubwuubwuubwu",
				"bwurrgbwurrgbwurrgbwurrgbwurrgbwurrgbwurrgbwurrgbwurrgbwurrgbwurrgbwurrgbwurrgbwurrgbwurrg",
				"brgrbrgrbrgrbrgrbrgrbrgrbrgrbrgrbrgrbrgrbrgrbrgrbrgrbrgrbrgrbrgrbrgrbrgrbrgrbrgrbrgr",
				"bbrgwbbbrgwbbbrgwbbbrgwbbbrgwbbbrgwbbbrgwbbbrgwbbbrgwbbbrgwbbbrgwbbbrgwbbbrgwbbbrgwbbbrgwbbbrgwb",
			},
			6,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.PossibleDesigns(tc.availablePatterns, tc.targetPatterns)
			if res != tc.possibleDesigns {
				t.Errorf("Expected %d, got %d", tc.possibleDesigns, res)
			}
		})
	}
}

func TestAllDesigns(t *testing.T) {
	testCases := []struct {
		availablePatterns []string
		targetPatterns    []string
		possibleDesigns   int
	}{
		{
			[]string{"a"},
			[]string{"aaa"},
			1,
		},
		{
			[]string{"a", "aa"},
			[]string{"aaa"},
			3,
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.AllDesigns(tc.availablePatterns, tc.targetPatterns)
			if res != tc.possibleDesigns {
				t.Errorf("Expected %d, got %d", tc.possibleDesigns, res)
			}
		})
	}
}
