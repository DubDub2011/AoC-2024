package solution

import gr "day12/grid"

func FencePricer(input [][]rune) int {
	grid := gr.New(input)
	perimiter, area := 0, 0
	for pos := range grid.Positions() {
		val := grid.GetPos(pos.X, pos.Y)

		up, right, down, left := pos.Up(), pos.Right(), pos.Down(), pos.Left()

		pointVal := grid.GetPos(up.X, up.Y)
		if pointVal != val {
			perimiter++
		}

		pointVal = grid.GetPos(right.X, right.Y)
		if pointVal != val {
			perimiter++
		}

		pointVal = grid.GetPos(down.X, down.Y)
		if pointVal != val {
			perimiter++
		}

		pointVal = grid.GetPos(left.X, left.Y)
		if pointVal != val {
			perimiter++
		}
		area++
	}
	return perimiter * area
}
