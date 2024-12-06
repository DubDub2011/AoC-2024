package solution

import (
	"day6/frame"
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
	moverFuncMap = map[int]func(frame.Position) frame.Position{
		Up:    frame.Position.Up,
		Right: frame.Position.Right,
		Down:  frame.Position.Down,
		Left:  frame.Position.Left,
	}
)

func RouteFinder(points [][]rune) (int, bool) {
	grid := frame.Grid{points}
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

func moveGuard(pos frame.Position, direction int, grid frame.Grid) (frame.Position, int, bool) {
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

func findGuard(grid frame.Grid) frame.Position {
	for x := 0; x < grid.GetLength(); x++ {
		for y := 0; y < grid.GetWidth(); y++ {
			if grid.GetPos(x, y) == GuardRune {
				return frame.Position{x, y}
			}
		}
	}
	panic("could not find guard")
}

func countVisitedPositions(grid frame.Grid) int {
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
