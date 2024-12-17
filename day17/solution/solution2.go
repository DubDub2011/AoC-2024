package solution

import (
	"fmt"
)

func Repeat(_, targetVals []int) {
	idx := 10000000000001
	for {
		program(idx, targetVals)
		if (idx-1)%80000000 == 0 {
			fmt.Printf("Got to %d, Closeness for boundary %v\n", idx, closeness)
			closeness = map[int]int{}
		}

		idx += 8
	}

}

var closeness = map[int]int{}

func program(startVal int, target []int) {
	val := startVal
	idx := len(target) - 1
	for {
		B := val % 8 // b is 0-7
		B = B ^ 2
		C := val / powInt(2, B) // 2 ^ (0-7)
		B = B ^ (C ^ 3)
		B = B % 8 // b is 0-7
		if target[idx] != B {
			closeness[idx]++
			return
		}
		val = val / 8
		if val == 0 {
			break
		}
		idx--
	}

	fmt.Printf("Part 2: %d", startVal)
}
