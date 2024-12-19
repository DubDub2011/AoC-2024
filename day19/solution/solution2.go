package solution

import "fmt"

func AllDesigns(availablePatterns []string, targetPatterns []string) int {
	possibleParts := make(map[string]struct{}, len(availablePatterns))
	for _, pattern := range availablePatterns {
		possibleParts[pattern] = struct{}{}
	}

	sum := 0
	for idx, targetPattern := range targetPatterns {
		fmt.Printf("Index: %d, Sum: %d\n", idx, sum)
		sum += getPossibilities(targetPattern, possibleParts)
	}

	return sum
}

func getPossibilities(targetPattern string, possibleParts map[string]struct{}) int {
	startingNode := Node{pattern: "", value: targetPattern}

	nodesToProcess := []Node{startingNode}
	sum := 0
	for {
		if len(nodesToProcess) == 0 {
			break
		}

		currNode := nodesToProcess[0]
		if currNode.value == "" {
			sum++
		}

		for part := range possibleParts {
			if len(currNode.value) < len(part) {
				continue
			}

			if currNode.value[:len(part)] == part {
				newNode := Node{pattern: part, value: currNode.value[len(part):]}
				currNode.children = append(currNode.children, newNode)
				nodesToProcess = append(nodesToProcess, newNode)
			}

		}

		nodesToProcess = nodesToProcess[1:]
	}

	return sum
}
