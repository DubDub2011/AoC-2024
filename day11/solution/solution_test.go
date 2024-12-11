package solution_test

import (
	"day11/solution"
	"fmt"
	"testing"
)

func TestStoneCount(t *testing.T) {
	testCases := []struct {
		input  []int
		count  int
		output int
	}{
		{[]int{0}, 1, 1}, // 1
		{[]int{0}, 2, 1}, // 2024
		{[]int{0}, 3, 2}, // 20 24
		{[]int{0}, 4, 4}, // 2 0 2 4
		{[]int{0}, 5, 4}, // 4048 0 4048 8096
		{[]int{0}, 6, 7}, // 40 48 0 40 48 80 96
		{[]int{125, 17}, 1, 3},
		{[]int{125, 17}, 2, 4},
		{[]int{125, 17}, 3, 5},
		{[]int{125, 17}, 4, 9},
		{[]int{125, 17}, 5, 13},
		{[]int{125, 17}, 6, 22},
		{[]int{125, 17}, 25, 55312},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.StoneCount(tc.input, tc.count)
			if res != tc.output {
				t.Errorf("Expected %d, got %d", tc.output, res)
			}
		})
	}
}

func TestStoneCountSliced(t *testing.T) {
	testCases := []struct {
		input  []int
		count  int
		output int
	}{
		{[]int{0}, 1, 1}, // 1
		{[]int{0}, 2, 1}, // 2024
		{[]int{0}, 3, 2}, // 20 24
		{[]int{0}, 4, 4}, // 2 0 2 4
		{[]int{0}, 5, 4}, // 4048 0 4048 8096
		{[]int{0}, 6, 7}, // 40 48 0 40 48 80 96
		{[]int{125, 17}, 1, 3},
		{[]int{125, 17}, 2, 4},
		{[]int{125, 17}, 3, 5},
		{[]int{125, 17}, 4, 9},
		{[]int{125, 17}, 5, 13},
		{[]int{125, 17}, 6, 22},
		{[]int{125, 17}, 25, 55312},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.StoneCountSliced(tc.input, tc.count)
			if res != tc.output {
				t.Errorf("Expected %d, got %d", tc.output, res)
			}
		})
	}
}

func BenchmarkStoneCount(b *testing.B) {
	count := 20
	input := []int{0, 5601550, 3914, 852, 50706, 68, 6, 645371}
	b.Run(fmt.Sprintf("Benchmark Normal %d", count), func(b *testing.B) {
		solution.StoneCount(input, count)
	})

	count += 5
	b.Run(fmt.Sprintf("Benchmark Normal %d", count), func(b *testing.B) {
		solution.StoneCount(input, count)
	})

	count += 5
	b.Run(fmt.Sprintf("Benchmark Normal %d", count), func(b *testing.B) {
		solution.StoneCount(input, count)
	})

	count += 5
	b.Run(fmt.Sprintf("Benchmark Normal %d", count), func(b *testing.B) {
		solution.StoneCount(input, count)
	})

}

func BenchmarkStoneCountSliced(b *testing.B) {
	count := 20
	input := []int{0, 5601550, 3914, 852, 50706, 68, 6, 645371}
	b.Run(fmt.Sprintf("Benchmark Slice %d", count), func(b *testing.B) {
		solution.StoneCountSliced(input, count)
	})

	count += 5
	b.Run(fmt.Sprintf("Benchmark Slice %d", count), func(b *testing.B) {
		solution.StoneCountSliced(input, count)
	})

	count += 5
	b.Run(fmt.Sprintf("Benchmark Slice %d", count), func(b *testing.B) {
		solution.StoneCountSliced(input, count)
	})

	count += 5
	b.Run(fmt.Sprintf("Benchmark Slice %d", count), func(b *testing.B) {
		solution.StoneCountSliced(input, count)
	})

}
