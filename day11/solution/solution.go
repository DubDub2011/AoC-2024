package solution

import "strconv"

func StoneCount(input []int, count int) int {
	res := make([]int64, len(input))
	for idx, val := range input {
		res[idx] = int64(val)
	}

	for idx := 0; idx < count; idx++ {
		res = applyRules(res)
	}

	return len(res)
}

func applyRules(input []int64) []int64 {
	res := make([]int64, 0, len(input))
	for _, val := range input {
		if val == 0 {
			res = append(res, 1)
			continue
		}

		strVal := strconv.FormatInt(val, 10)
		if len(strVal)%2 == 0 {
			left, right := splitVal(strVal)
			res = append(res, left)
			res = append(res, right)
			continue
		}

		res = append(res, val*2024)
	}
	return res
}

func splitVal(inp string) (int64, int64) {
	leftSide, rightSide := inp[0:(len(inp)/2)], inp[(len(inp)/2):len(inp)]
	left, err := strconv.ParseInt(leftSide, 10, 0)
	if err != nil {
		panic(err)
	}

	right, err := strconv.ParseInt(rightSide, 10, 0)
	if err != nil {
		panic(err)
	}

	return left, right
}
