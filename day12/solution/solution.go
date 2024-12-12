package solution

import gr "day12/grid"

func FencePricer(input [][]rune) int {
	grid := gr.New(input)
	perimiter, area := 0, 0
	for pos := range grid.Positions() {
		val := grid.GetPos(pos.X, pos.Y)

		up, right, down, left := pos.Up(), pos.Right(), pos.Down(), pos.Left()
		if grid.GetPos(up.X, up.Y) != val {
			perimiter++
		}
		if grid.GetPos(right.X, right.Y) != val {
			perimiter++
		}
		if grid.GetPos(down.X, down.Y) != val {
			perimiter++
		}
		if grid.GetPos(left.X, left.Y) != val {
			perimiter++
		}
		area++
	}
	return perimiter * area
}
