package grid

import (
	"iter"
)

type Grid struct {
	points [][]rune
}

func New(points [][]rune) Grid {
	return Grid{points}
}

func (g Grid) GetPos(x, y int) rune {
	if y < 0 || y >= len(g.points[0]) {
		return ' '
	}

	if x < 0 || x >= len(g.points) {
		return ' '
	}

	return g.points[y][x]
}

func (g Grid) SetPos(x, y int, val rune) {
	if x < 0 || x >= len(g.points[0]) {
		return
	}

	if y < 0 || y >= len(g.points) {
		return
	}

	g.points[y][x] = val
}

func (g Grid) GetWidth() int {
	return len(g.points[0])
}

func (g Grid) GetLength() int {
	return len(g.points)
}

func (g Grid) String() string {
	res := "\n"
	for _, row := range g.points {
		res += string(row)
		res += "\n"
	}
	return res
}

func (g Grid) Positions() iter.Seq[Position] {
	return func(yield func(Position) bool) {
		for row := range g.points {
			for col := range g.points[0] {
				if !yield(Position{row, col}) {
					return
				}
			}
		}
	}
}
