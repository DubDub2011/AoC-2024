package solution

import (
	gr "day8/grid"
)

const (
	EmptyRune = '.'
)

func Resonance(input [][]rune) int {
	grid := gr.New(input)
	total := 0
	for pos := range grid.Positions() {
		val := grid.GetPos(pos.X, pos.Y)
		if val != EmptyRune {
			total++
		}
	}
	total -= 1
	return total
}
