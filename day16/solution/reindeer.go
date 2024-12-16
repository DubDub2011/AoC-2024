package solution

import gr "day16/grid"

var directionMovers = map[int]func(gr.Position) gr.Position{
	0: gr.Position.Right,
	1: gr.Position.Down,
	2: gr.Position.Left,
	3: gr.Position.Up,
}

type Reindeer struct {
	score     int
	direction int
	currPos   gr.Position
	path      []gr.Position
}

func NewReindeer(startPos gr.Position) *Reindeer {
	return &Reindeer{
		0,
		0,
		startPos,
		[]gr.Position{startPos},
	}
}

func (rndr *Reindeer) Clone() *Reindeer {
	pathClone := make([]gr.Position, len(rndr.path))
	copy(pathClone, rndr.path)
	return &Reindeer{
		rndr.score,
		rndr.direction,
		rndr.currPos,
		pathClone,
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

func (rndr *Reindeer) Move() {
	rndr.currPos = directionMovers[rndr.direction](rndr.currPos)
	rndr.path = append(rndr.path, rndr.currPos)
	rndr.score += 1
}

func (rndr *Reindeer) CanTurn(grid gr.Grid) (bool, bool) {
	left, right := (rndr.direction+3)%4, (rndr.direction+1)%4
	leftPos := directionMovers[left](rndr.currPos)
	rightPos := directionMovers[right](rndr.currPos)
	return grid.GetPos(leftPos.X, leftPos.Y) != '#', grid.GetPos(rightPos.X, rightPos.Y) != '#'
}
