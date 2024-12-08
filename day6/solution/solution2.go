package solution

import (
	gr "day6/grid"
)

const (
	UpRune    = '^'
	RightRune = '>'
	DownRune  = 'v'
	LeftRune  = '<'
) // Direction runes

var (
	directionRuneMap = map[int]rune{
		Up:    UpRune,
		Right: RightRune,
		Down:  DownRune,
		Left:  LeftRune,
	}
)

func RouteBreaker(points [][]rune) int {
	grid := gr.Grid{points}
	guardPos := findGuard(grid)

	total := 0
	for x := 0; x < grid.GetLength(); x++ {
		for y := 0; y < grid.GetWidth(); y++ {
			if grid.GetPos(x, y) == WallRune {
				continue
			}

			newGrid := copyGrid(grid)
			newGrid.SetPos(x, y, WallRune)

			if checkRoute(guardPos, newGrid) {
				total++
			}
		}
	}
	return total
}

func copyGrid(grid gr.Grid) gr.Grid {
	newPoints := make([][]rune, grid.GetLength())
	for idx := 0; idx < grid.GetLength(); idx++ {
		newPoints[idx] = make([]rune, grid.GetWidth())
		copy(newPoints[idx], grid.Points[idx])
	}
	return gr.Grid{newPoints}
}

func checkRoute(pos gr.Position, grid gr.Grid) bool {
	direction := Up
	grid.SetPos(pos.X, pos.Y, directionRuneMap[direction]) // set current position to already visited

	for {
		targetPosition := moverFuncMap[direction](pos)
		targetVal := grid.GetPos(targetPosition.X, targetPosition.Y)
		if targetVal == ' ' {
			return false
		}

		if targetVal == WallRune {
			direction = (direction + 1) % 4
			continue
		}

		if grid.GetPos(targetPosition.X, targetPosition.Y) == directionRuneMap[direction] {
			return true
		}

		pos = targetPosition
		grid.SetPos(pos.X, pos.Y, directionRuneMap[direction])
	}
}
