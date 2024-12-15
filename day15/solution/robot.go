package solution

import gr "day15/grid"

func NewRobot(startPos gr.Position, commands []rune) *Robot {
	return &Robot{
		startPos,
		commands,
	}
}

type Robot struct {
	gr.Position
	commands []rune
}

func (r *Robot) GetMoveDirection() func(gr.Position) gr.Position {
	if len(r.commands) == 0 {
		return nil
	}
	inp := r.commands[0]
	r.commands = r.commands[1:]

	if inp == 'v' {
		return gr.Position.Down
	} else if inp == '>' {
		return gr.Position.Right
	} else if inp == '^' {
		return gr.Position.Up
	} else if inp == '<' {
		return gr.Position.Left
	} else {
		panic("unknown direction")
	}
}

func (r *Robot) Move(mover func(gr.Position) gr.Position) {
	r.Position = mover(r.Position)
}
