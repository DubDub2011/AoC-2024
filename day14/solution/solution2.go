package solution

import (
	gr "day14/grid"
	"fmt"
	"strings"
)

func RobotChristmasTree(robotData [][]int) int {
	// the puzzle here is so vague... going to experiment for now
	robots := []*Robot{}
	for _, robot := range robotData {
		robots = append(robots, &Robot{gr.Position{robot[0], robot[1]}, robot[2], robot[3]})
	}

	count := 0
	for {
		robotPositions := make(map[gr.Position]bool, len(robots))
		for _, robot := range robots {
			robot.Move(1)
			robotPositions[robot.Position()] = true
		}

		if robotPositions[gr.Position{WIDTH / 2, 0}] &&
			robotPositions[gr.Position{WIDTH/2 - 1, 1}] &&
			robotPositions[gr.Position{WIDTH/2 + 1, 1}] &&
			robotPositions[gr.Position{WIDTH/2 - 2, 2}] &&
			robotPositions[gr.Position{WIDTH/2 + 2, 2}] &&
			robotPositions[gr.Position{WIDTH/2 - 3, 3}] &&
			robotPositions[gr.Position{WIDTH/2 + 3, 3}] &&
			robotPositions[gr.Position{WIDTH/2 - 4, 4}] &&
			robotPositions[gr.Position{WIDTH/2 + 4, 4}] &&
			robotPositions[gr.Position{WIDTH/2 - 5, 5}] &&
			robotPositions[gr.Position{WIDTH/2 + 5, 5}] { // this is lazy but might be good enough
			draw(robots)
			break
		}
		count++

		if count%100000 == 0 {
			fmt.Printf("Count: %d\n", count)
		}
	}

	return count
}

func draw(robots []*Robot) {
	runes := make([][]rune, LENGTH)
	for idx := range runes {
		runes[idx] = []rune(strings.Repeat(".", WIDTH))
	}
	grid := gr.New(runes)
	for _, robot := range robots {
		gridPos := grid.GetPos(robot.pos.X, robot.pos.Y)
		if gridPos == '.' {
			grid.SetPos(robot.pos.X, robot.pos.Y, 'A')
		} else {
			grid.SetPos(robot.pos.X, robot.pos.Y, rune(gridPos+1))
		}
	}
	fmt.Printf("Grid: %s\n", grid.String())
}
