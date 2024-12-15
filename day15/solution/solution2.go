package solution

import (
	gr "day15/grid"
)

func BigBoxChecker(mapData [][]rune, commands []rune) int {
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
		mover, isHoz := robot.GetMoveDirection()
		if mover == nil {
			break
		}

		var positions []gr.Position
		if isHoz {
			positions = getValidHozMovePositionsBigBox(robot.Position, grid, mover)
		} else {
			positions = getValidVerMovePositionsBigBox(robot.Position, grid, mover)
		}

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
		if val == '[' {
			sum += pos.X + pos.Y*100
		}
	}

	return sum
}

func getValidHozMovePositionsBigBox(pos gr.Position, grid gr.Grid, mover func(gr.Position) gr.Position) []gr.Position {
	res := []gr.Position{pos}

	for {
		pos = mover(pos)
		if grid.GetPos(pos.X, pos.Y) == '#' {
			return []gr.Position{}
		} else if grid.GetPos(pos.X, pos.Y) == '[' {
			res = append(res, pos)
			pos = mover(pos)
			res = append(res, pos)
		} else if grid.GetPos(pos.X, pos.Y) == ']' {
			res = append(res, pos)
			pos = mover(pos)
			res = append(res, pos)
		} else if grid.GetPos(pos.X, pos.Y) == '.' {
			break
		}
	}

	return res
}

func getValidVerMovePositionsBigBox(pos gr.Position, grid gr.Grid, mover func(gr.Position) gr.Position) []gr.Position {
	res := []gr.Position{pos}

	edgePositions := []gr.Position{pos}
	for {
		if len(edgePositions) == 0 {
			break
		}

		newEdgePositions := []gr.Position{}
		for _, edgePos := range edgePositions {
			edgePos = mover(edgePos)
			if grid.GetPos(edgePos.X, edgePos.Y) == '#' {
				return []gr.Position{}
			} else if grid.GetPos(edgePos.X, edgePos.Y) == '[' {
				res = append(res, edgePos)
				newEdgePositions = append(newEdgePositions, edgePos)
				res = append(res, edgePos.Right())
				newEdgePositions = append(newEdgePositions, edgePos.Right())
			} else if grid.GetPos(edgePos.X, edgePos.Y) == ']' {
				res = append(res, edgePos)
				newEdgePositions = append(newEdgePositions, edgePos)
				res = append(res, edgePos.Left())
				newEdgePositions = append(newEdgePositions, edgePos.Left())
			} else if grid.GetPos(edgePos.X, edgePos.Y) == '.' {
				continue
			}
		}

		edgePositions = newEdgePositions
	}

	allKeys := make(map[gr.Position]bool)
	list := []gr.Position{}
	for _, pos := range res {
		if !allKeys[pos] {
			allKeys[pos] = true
			list = append(list, pos)
		}
	}

	return list
}
