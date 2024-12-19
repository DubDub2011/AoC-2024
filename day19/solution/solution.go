package solution

import "fmt"

func PossibleDesigns(availablePatterns []string, targetPatterns []string) int {
	possibleParts := make(map[string]struct{}, len(availablePatterns))
	for _, pattern := range availablePatterns {
		possibleParts[pattern] = struct{}{}
	}

	for part := range possibleParts {
		for idx := 0; idx < len(part); idx++ {
			left, right := part[:idx], part[idx:]
			_, leftExists := possibleParts[left]
			_, rightExists := possibleParts[right]
			if leftExists && rightExists {
				delete(possibleParts, part)
				break
			}
		}
	}

	sum := 0
	for idx, targetPattern := range targetPatterns {
		fmt.Printf("Index: %d, Sum: %d\n", idx, sum)
		if checkPossible(targetPattern, possibleParts) {
			sum++
		}
	}

	return sum
}

func checkPossible(targetPattern string, possibleParts map[string]struct{}) bool {
	startingNode := Node{pattern: "", value: targetPattern}
	nodes := []Node{startingNode}
	for {
		if len(nodes) == 0 {
			break
		}

		currNode := nodes[0]
		if currNode.value == "" {
			return true
		}

		for part := range possibleParts {
			if len(currNode.value) < len(part) {
				continue
			}

			if currNode.value[:len(part)] == part {
				newNode := Node{pattern: part, value: currNode.value[len(part):]}
				nodes = append(nodes, newNode)
			}

		}

		nodes = nodes[1:]
	}

	return false
}

type Node struct {
	children []Node
	pattern  string
	value    string
}
