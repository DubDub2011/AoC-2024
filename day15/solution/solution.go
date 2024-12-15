package solution

import gr "day15/grid"

func BoxChecker(mapData [][]rune, commands []rune) int {
	grid := gr.New(mapData)
	startPos := gr.Position{X: grid.GetLength() / 2, Y: grid.GetWidth() / 2}

	robot := NewRobot(startPos, commands)

	print(robot)

	return 0
}
