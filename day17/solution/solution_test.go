package solution_test

import (
	"day17/solution"
	"fmt"
	"testing"
)

func TestThreeBitComputer(t *testing.T) {
	tests := []struct {
		registers []int
		program   []int
		output    string
	}{
		{
			[]int{729, 0, 0},
			[]int{0, 1, 5, 4, 3, 0},
			"4,6,3,5,6,3,5,2,1,0",
		},
	}

	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test%d", idx), func(t *testing.T) {
			if got := solution.ThreeBitComputer(test.registers, test.program); got != test.output {
				t.Errorf("ThreeBitComputer() = %v, want %v", got, test.output)
			}
		})
	}
}
