package solution

import (
	gr "day6/grid"
)

const (
	GuardRune          = '^'
	WallRune           = '#'
	AlreadyVisitedRune = 'X'
) // runes

const (
	Up    = iota
	Right = iota
	Down  = iota
	Left  = iota
) // directions

var (
	moverFuncMap = map[int]func(gr.Position) gr.Position{
		Up:    gr.Position.Up,
		Right: gr.Position.Right,
		Down:  gr.Position.Down,
		Left:  gr.Position.Left,
	}
)

func RouteFinder(points [][]rune) (int, bool) {
	grid := gr.Grid{points}
	guardPos := findGuard(grid)
	currDir := Up

	for {
		var finished bool
		guardPos, currDir, finished = moveGuard(guardPos, currDir, grid)
		if finished {
			break
		}
	}

	return countVisitedPositions(grid), true
}

func moveGuard(pos gr.Position, direction int, grid gr.Grid) (gr.Position, int, bool) {
	grid.SetPos(pos.X, pos.Y, AlreadyVisitedRune) // set current position to already visited

	for {
		targetPosition := moverFuncMap[direction](pos)
		targetVal := grid.GetPos(targetPosition.X, targetPosition.Y)
		if targetVal == ' ' {
			return targetPosition, direction, true
		}

		if targetVal == WallRune {
			direction = (direction + 1) % 4
			continue
		}

		grid.SetPos(targetPosition.X, targetPosition.Y, GuardRune)
		return targetPosition, direction, false
	}
}

func findGuard(grid gr.Grid) gr.Position {
	for x := 0; x < grid.GetLength(); x++ {
		for y := 0; y < grid.GetWidth(); y++ {
			if grid.GetPos(x, y) == GuardRune {
				return gr.Position{x, y}
			}
		}
	}
	panic("could not find guard")
}

func countVisitedPositions(grid gr.Grid) int {
	total := 0
	for x := 0; x < grid.GetLength(); x++ {
		for y := 0; y < grid.GetWidth(); y++ {
			if grid.GetPos(x, y) == AlreadyVisitedRune {
				total++
			}
		}
	}
	return total
}
