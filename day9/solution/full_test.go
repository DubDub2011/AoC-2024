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
		{[]byte("20"), 0},                     // 00
		{[]byte("122"), 3},                    // 0..11
		{[]byte("12222"), 13},                 // 0..11..22
		{[]byte("2333133121414131402"), 1928}, // from task
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

func TestDiskFragmenter2(t *testing.T) {
	testCases := []struct {
		input    []byte
		expected int
	}{
		{[]byte("20"), 0},              // 00
		{[]byte("122"), 3},             // 0..11
		{[]byte("12222"), 13},          // 0..11..22
		{[]byte("01010101101"), 0},     // ....01
		{[]byte("010101011111111"), 4}, // ....0.1.2.3
		{[]byte("1234554321"), 234},
		{[]byte("112345678"), 1446},
		{[]byte("00122002200"), 13}, // 0..11..22
		{[]byte("1313165"), 169},    // 0..11..22

		{[]byte("2333133121414131402"), 2858}, // from task
	}

	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx), func(t *testing.T) {
			res := solution.DiskFragmenter2(testCase.input)
			if res != testCase.expected {
				t.Errorf("expected %d, got %d", testCase.expected, res)
			}
		})
	}
}
