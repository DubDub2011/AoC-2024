package solution

import gr "day10/grid"

func HeightmapRouter(heightmap [][]rune) int {
	grid := gr.New(heightmap)
	trailheads := getTrailheads(grid)

	sum := 0
	for _, pos := range trailheads {
		sum += route(grid, pos)
	}

	return sum
}

func route(grid gr.Grid, startingPos gr.Position) int {
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

	// remove duplicates
	currPositions = removeDuplicates(currPositions)
	return len(currPositions)
}

func removeDuplicates(slice []gr.Position) []gr.Position {
	allKeys := make(map[gr.Position]bool)
	list := []gr.Position{}
	for _, item := range slice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func getTrailheads(grid gr.Grid) []gr.Position {
	trailheads := make([]gr.Position, 0)
	for pos := range grid.Positions() {
		if grid.GetPos(pos.X, pos.Y) == '0' {
			trailheads = append(trailheads, pos)
		}
	}
	return trailheads
}

func runeToInt(val rune) int {
	if val < '0' || val > '9' {
		return 0
	}
	return int(val) - 48
}

func intToRune(val int) rune {
	if val > 9 {
		panic("can't be greater than 10")
	}
	if val < 0 {
		panic("can't be negative")
	}
	return rune(val + 48)
}
