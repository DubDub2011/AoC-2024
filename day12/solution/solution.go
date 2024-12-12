package solution

import gr "day12/grid"

func FencePricer(input [][]rune) int {
	grid := gr.New(input)

	alreadyVisited := make(map[gr.Position]bool)
	total := 0
	for pos := range grid.Positions() {
		if alreadyVisited[pos] {
			continue
		}

		regionPositions := []gr.Position{pos}
		val := grid.GetPos(pos.X, pos.Y)
		perimiter, area := 0, 0
		for idx := 0; idx < len(regionPositions); idx++ {
			regionPos := regionPositions[idx]
			if alreadyVisited[regionPos] {
				continue
			}

			up, right, down, left := regionPos.Up(), regionPos.Right(), regionPos.Down(), regionPos.Left()
			if grid.GetPos(up.X, up.Y) == val {
				regionPositions = append(regionPositions, up)
			} else {
				perimiter++
			}

			if grid.GetPos(right.X, right.Y) == val {
				regionPositions = append(regionPositions, right)
			} else {
				perimiter++
			}

			if grid.GetPos(down.X, down.Y) == val {
				regionPositions = append(regionPositions, down)
			} else {
				perimiter++
			}

			if grid.GetPos(left.X, left.Y) == val {
				regionPositions = append(regionPositions, left)
			} else {
				perimiter++
			}
			area++
			alreadyVisited[regionPos] = true
		}

		total += area * perimiter
	}
	return total
}
