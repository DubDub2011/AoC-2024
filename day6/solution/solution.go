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

var currDir = Up // bit lazy...

func RouteFinder(points [][]rune) int {
	grid := frame.Grid{points}
	guardPos := findGuard(grid)
	currDir = Up

	for {
		var finished bool
		guardPos, finished = moveGuard(guardPos, grid)
		if finished {
			break
		}
	}

	return countVisitedPositions(grid)
}

func moveGuard(pos frame.Position, grid frame.Grid) (frame.Position, bool) {
	grid.SetPos(pos.X, pos.Y, AlreadyVisitedRune) // set current position to already visited

	for {
		targetPosition := moverFuncMap[currDir](pos)
		targetVal := grid.GetPos(targetPosition.X, targetPosition.Y)
		if targetVal == ' ' {
			return targetPosition, true
		}

		if targetVal == WallRune {
			currDir = (currDir + 1) % 4
			continue // might get stuck in here if we somehow get boxed in, shouldn't happen though
		}

		grid.SetPos(targetPosition.X, targetPosition.Y, GuardRune)
		return targetPosition, false
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
