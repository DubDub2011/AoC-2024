package solution

import (
	gr "day15/grid"
)

func BoxChecker(mapData [][]rune, commands []rune) int {
	grid := gr.New(mapData)
	var startPos gr.Position
	for pos := range grid.Positions() {
		if grid.GetPos(pos.X, pos.Y) == '@' {
			startPos = pos
			break
		}
	}

	robot := NewRobot(startPos, commands)

	for {
		mover, _ := robot.GetMoveDirection()
		if mover == nil {
			break
		}

		positions := getValidMovePositions(robot.Position, grid, mover)
		for idx := len(positions) - 1; idx >= 0; idx-- {
			pos := positions[idx]
			val := grid.GetPos(pos.X, pos.Y)
			grid.SetPos(pos.X, pos.Y, '.')
			pos = mover(pos)
			grid.SetPos(pos.X, pos.Y, val)
		}

		if len(positions) != 0 {
			robot.Position = mover(robot.Position)
		}
	}

	sum := 0
	for pos := range grid.Positions() {
		val := grid.GetPos(pos.X, pos.Y)
		if val == 'O' {
			sum += pos.X + pos.Y*100
		}
	}

	return sum
}

func getValidMovePositions(pos gr.Position, grid gr.Grid, mover func(gr.Position) gr.Position) []gr.Position {
	res := []gr.Position{pos}
	for {
		pos = mover(pos)
		if grid.GetPos(pos.X, pos.Y) == '#' {
			return []gr.Position{}
		} else if grid.GetPos(pos.X, pos.Y) == 'O' {
			res = append(res, pos)
		} else if grid.GetPos(pos.X, pos.Y) == '.' {
			break
		}
	}

	return res
}
