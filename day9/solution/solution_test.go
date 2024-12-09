package solution_test

import (
	"day9/solution"
	"fmt"
	"testing"
)

func TestDiskFragmenter(t *testing.T) {
	testCases := []struct {
		input    []byte
		expected int
	}{
		{[]byte("20"), 0},
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.DiskFragmenter(testCase.input)
			if res != testCase.expected {
				t.Errorf("expected %d, got %d", testCase.expected, res)
			}
		})
	}
}
