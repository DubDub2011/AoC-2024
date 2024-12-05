package solution

type grid struct {
	runes [][]rune
}

func (g grid) getPos(x, y int) rune {
	if x < 0 || x >= len(g.runes[0]) {
		return ' '
	}

	if y < 0 || y >= len(g.runes) {
		return ' '
	}

	return g.runes[y][x]
}

func (g grid) getWidth() int {
	return len(g.runes[0])
}

func (g grid) getLength() int {
	return len(g.runes)
}
