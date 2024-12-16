package solution

import gr "day16/grid"

var directionMovers = map[int]func(gr.Position) gr.Position{
	0: gr.Position.Right,
	1: gr.Position.Down,
	2: gr.Position.Left,
	3: gr.Position.Up,
}

type Reindeer struct {
	score       int
	direction   int
	currPos     gr.Position
	previousPos gr.Position
	done        bool
}

func NewReindeer(startPos gr.Position) *Reindeer {
	return &Reindeer{
		0,
		0,
		startPos,
		startPos,
		false,
	}
}

func (rndr *Reindeer) Clone() *Reindeer {
	return &Reindeer{
		rndr.score,
		rndr.direction,
		rndr.currPos,
		rndr.previousPos,
		rndr.done,
	}
}

func (rndr *Reindeer) TurnLeft() {
	rndr.direction = (rndr.direction + 3) % 4
	rndr.score += 1000
}

func (rndr *Reindeer) TurnRight() {
	rndr.direction = (rndr.direction + 1) % 4
	rndr.score += 1000

}

func (rndr *Reindeer) Move() bool {
	if rndr.done {
		return false
	}
	rndr.previousPos = rndr.currPos
	rndr.currPos = directionMovers[rndr.direction](rndr.currPos)
	rndr.score += 1

	return true
}

func (rndr *Reindeer) CanTurn(grid gr.Grid) (bool, bool) {
	left, right := (rndr.direction+3)%4, (rndr.direction+1)%4
	leftPos := directionMovers[left](rndr.currPos)
	rightPos := directionMovers[right](rndr.currPos)
	return grid.GetPos(leftPos.X, leftPos.Y) == '#', grid.GetPos(rightPos.X, rightPos.Y) == '#'
}
