package solution

type position struct {
	x int
	y int
}

func (p position) right() position {
	return position{
		x: p.x + 1,
		y: p.y,
	}
}

func (p position) left() position {
	return position{
		x: p.x - 1,
		y: p.y,
	}
}

func (p position) up() position {
	return position{
		x: p.x,
		y: p.y - 1,
	}
}

func (p position) down() position {
	return position{
		x: p.x,
		y: p.y + 1,
	}
}

func (p position) upRight() position {
	return p.up().right()
}

func (p position) upLeft() position {
	return p.up().left()
}

func (p position) downRight() position {
	return p.down().right()
}

func (p position) downLeft() position {
	return p.down().left()
}
