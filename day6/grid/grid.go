package grid

type Grid struct {
	Points [][]rune
}

func (g Grid) GetPos(x, y int) rune {
	if y < 0 || y >= len(g.Points[0]) {
		return ' '
	}

	if x < 0 || x >= len(g.Points) {
		return ' '
	}

	return g.Points[y][x]
}

func (g Grid) SetPos(x, y int, val rune) {
	if x < 0 || x >= len(g.Points[0]) {
		panic("set invalid position")
	}

	if y < 0 || y >= len(g.Points) {
		panic("set invalid position")
	}

	g.Points[y][x] = val
}

func (g Grid) GetWidth() int {
	return len(g.Points[0])
}

func (g Grid) GetLength() int {
	return len(g.Points)
}

func (g Grid) String() string {
	res := "\n"
	for _, row := range g.Points {
		res += string(row)
		res += "\n"
	}
	return res
}
