package solution

import (
	gr "day14/grid"
)

const (
	WIDTH  = 101
	LENGTH = 103
)

func RobotPredictor(robotData [][]int, seconds int) int {
	robots := []*Robot{}
	for _, robot := range robotData {
		robots = append(robots, &Robot{gr.Position{robot[0], robot[1]}, robot[2], robot[3]})
	}

	for _, robot := range robots {
		robot.Move(seconds)
	}

	// X is distance from the left, Y is distance from the top
	topLeft, topRight, bottomLeft, bottomRight := 0, 0, 0, 0
	for _, robot := range robots {
		if robot.pos.X < WIDTH/2 && robot.pos.Y < LENGTH/2 {
			topLeft++
		} else if robot.pos.X < WIDTH/2 && robot.pos.Y > LENGTH/2 {
			bottomLeft++
		} else if robot.pos.X > WIDTH/2 && robot.pos.Y < LENGTH/2 {
			topRight++
		} else if robot.pos.X > WIDTH/2 && robot.pos.Y > LENGTH/2 {
			bottomRight++
		}
	}

	return topLeft * topRight * bottomLeft * bottomRight
}

type Robot struct {
	pos    gr.Position
	vx, vy int
}

func (r *Robot) Position() gr.Position {
	return r.pos
}

func (r *Robot) Move(amount int) {
	r.pos.X += r.vx * amount
	r.pos.X = r.pos.X % WIDTH
	if r.pos.X < 0 {
		r.pos.X += WIDTH
	}

	r.pos.Y += r.vy * amount
	r.pos.Y = r.pos.Y % LENGTH
	if r.pos.Y < 0 {
		r.pos.Y += LENGTH
	}
}
