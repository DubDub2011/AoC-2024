package solution

func PagePrint(orderingRules [][]int, printUpdates [][]int) int {
	orderRules := make(map[int][]int, 0)
	for _, pair := range orderingRules {
		orderRules[pair[0]] = append(orderRules[pair[0]], pair[1])
	}

	total := 0
	for _, update := range printUpdates {
		previousValues := make([]int, 0)
		updateFailed := false
		for _, val := range update {
			unacceptablePriors := orderRules[val]
			if !valueAcceptable(unacceptablePriors, previousValues) {
				updateFailed = true
				break
			}

			previousValues = append(previousValues, val)
		}

		if updateFailed {
			continue
		}

		total += update[len(update)/2]
	}
	return total
}

func PagePrint2(orderingRules [][]int, printUpdates [][]int) int {
	orderRules := make(map[int][]int, 0)
	for _, pair := range orderingRules {
		orderRules[pair[0]] = append(orderRules[pair[0]], pair[1])
	}

	total := 0
	for _, update := range printUpdates {
		failedUpdate := false
		sorted := make([]int, 0, len(update))
		for idx := len(update) - 1; idx >= 0; {
			val := update[idx]

			previousValues := update[:idx]
			unacceptablePriors := orderRules[val]
			if !valueAcceptable(unacceptablePriors, previousValues) {
				failedUpdate = true
				update = append([]int{val}, update[:idx]...) // move value to the front
				continue
			} else {
				sorted = append([]int{val}, sorted...) // value now in correct place, move to next
				idx--
			}
		}

		if !failedUpdate { // we only want failed updates
			continue
		}

		middleIndex := len(sorted) / 2
		total += sorted[middleIndex]
	}
	return total

}

func valueAcceptable(unacceptableValues, priorValues []int) bool {
	for _, prevValue := range priorValues {
		for _, unacceptableValue := range unacceptableValues {
			if prevValue == unacceptableValue {
				return false
			}
		}
	}
	return true
}
