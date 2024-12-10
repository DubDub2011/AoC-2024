package solution

import gr "day10/grid"

func HeightmapUniquePaths(heightmap [][]rune) int {
	grid := gr.New(heightmap)
	trailheads := getTrailheads(grid)

	sum := 0
	for _, pos := range trailheads {
		sum += route2(grid, pos)
	}

	return sum
}

func route2(grid gr.Grid, startingPos gr.Position) int {
	currPositions := []gr.Position{startingPos}
	for target := 1; target < 10; target++ {
		newPositions := make([]gr.Position, 0)

		for _, pos := range currPositions {
			if up := pos.Up(); grid.GetPos(up.X, up.Y) == intToRune(target) {
				newPositions = append(newPositions, up)
			}
			if down := pos.Down(); grid.GetPos(down.X, down.Y) == intToRune(target) {
				newPositions = append(newPositions, down)
			}

			if left := pos.Left(); grid.GetPos(left.X, left.Y) == intToRune(target) {
				newPositions = append(newPositions, left)
			}

			if right := pos.Right(); grid.GetPos(right.X, right.Y) == intToRune(target) {
				newPositions = append(newPositions, right)
			}
		}
		currPositions = newPositions
		if len(currPositions) == 0 {
			break
		}
	}

	return len(currPositions)
}
