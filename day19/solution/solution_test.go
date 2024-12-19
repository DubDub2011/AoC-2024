package solution_test

import (
	"day19/solution"
	"fmt"
	"testing"
)

func TestPatternBuilder(t *testing.T) {
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
